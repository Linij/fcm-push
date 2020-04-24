#!/bin/bash
if [ ${CI_COMMIT_REF_NAME} == "develop" ]
then
    echo [检测到开发环境,POD数量变更为1个]
    replicas=1
    echo [检测到开发环境,NFS服务端为${NFS_DEV}:${NFS_DEV_DIR}]
    NFS_online=${NFS_DEV}
    NFS_online_DIR=${NFS_DEV_DIR}
    echo [检测到开发环境,KUBE服务端为${KUBE_DEV}:${KUBE_DEV_PORT}]
    KUBE_online=${KUBE_DEV}
    KUBE_online_PORT=${KUBE_DEV_PORT}
    Labels_Selector="debug"
else
    Labels_Selector="high"
fi

containersName=${ProjectName}-${CI_JOB_ID}
function PrintPOD()
{
cat<<SUB
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: ${ProjectName}
spec:
  replicas: ${replicas}
  template:
    metadata:
      labels:
        name: ${ProjectName}
    spec:
      containers:
      - name: ${containersName}
        image: ${IMAGE}
        securityContext:
          capabilities:
            add:
            - SYS_PTRACE
            - SYSLOG
        imagePullPolicy: Always
        volumeMounts:
        - name: localtime
          mountPath: /etc/localtime
        - name: htdocs
          mountPath: ${ProjecHtdocsDir}
        - name: super
          mountPath:  /super_config
        command : ["/bin/bash","/start.sh"]
      volumes:
      - name: localtime
        hostPath:
            path: /etc/localtime
      - name: htdocs
        nfs:
            server: ${NFS_online}
            path: ${NFS_online_DIR}/${ProjectName}
       - name: super
        nfs:
            server: ${NFS_online}
            path: ${NFS_online_DIR}/${ProjectName}/_super_${CI_COMMIT_REF_NAME}
      restartPolicy: Always
      nodeSelector:
        performance: ${Labels_Selector}

SUB
}

function RollUpdate()
{
    PrintPOD | kubectl -s ${KUBE_online}:${KUBE_online_PORT} apply -f -
}



find ${ProjectDir}/.gitlab-ci.yml -type f -print | xargs md5sum > /tmp/${ProjectName}_new.md5 
find ${ProjectDir}/_gitlabci/deploy.sh -type f -print | xargs md5sum >> /tmp/${ProjectName}_new.md5
find ${ProjectDir}/_gitlabci/BuildDocker.sh -type f -print | xargs md5sum >> /tmp/${ProjectName}_new.md5
find ${ProjectDir}/_gitlabci/start.sh -type f -print | xargs md5sum >> /tmp/${ProjectName}_new.md5


if [ ! -e "/tmp/${ProjectName}_old.md5" ]
then
    echo [检测到初次生成MD5文件，准备滚动更新]
    RollUpdate
else
    if [ -z "`diff /tmp/${ProjectName}_new.md5 /tmp/${ProjectName}_old.md5`" ]
    then
        echo [检测到文件无变化，不更新]
    else
        echo [检测到文件有变化，准备滚动更新]
        RollUpdate
    fi
fi
cp -f /tmp/${ProjectName}_new.md5 /tmp/${ProjectName}_old.md5
