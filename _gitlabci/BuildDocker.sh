#!/bin/bash
if [ "$CI_COMMIT_REF_NAME" == "online" ]
then
    CI_COMMIT_REF_NAME="online"
else
    CI_COMMIT_REF_NAME="develop"
fi


function BuildDockerImage()
{
    docker build --pull -t $IMAGE --build-arg CI_COMMIT_REF_NAME=${CI_COMMIT_REF_NAME} --build-arg ProjecHtdocsDir=${ProjecHtdocsDir} -f ${ProjectDir}/_gitlabci/Dockerfile ${ProjectDir}
    docker push $IMAGE
}



find ${ProjectDir}/.gitlab-ci.yml -type f -print | xargs md5sum > /tmp/${ProjectName}_buid_docker_new.md5 
find ${ProjectDir}/_gitlabci/BuildDocker.sh -type f -print | xargs md5sum >> /tmp/${ProjectName}_buid_docker_new.md5
find ${ProjectDir}/_gitlabci/Dockerfile -type f -print | xargs md5sum >> /tmp/${ProjectName}_buid_docker_new.md5
find ${ProjectDir}/_gitlabci/start.sh -type f -print | xargs md5sum >> /tmp/${ProjectName}_buid_docker_new.md5


if [ ! -e "/tmp/${ProjectName}_buid_docker_old.md5" ]
then
    echo [检测到初次生成MD5文件，准备构建镜像]
    BuildDockerImage
else
    if [ -z "`diff /tmp/${ProjectName}_buid_docker_new.md5 /tmp/${ProjectName}_buid_docker_old.md5`" ]
    then
        echo [检测到文件无变化，不构建镜像]
    else
        echo [检测到文件有变化，准备构建镜像]
        BuildDockerImage
    fi
fi
cp -f /tmp/${ProjectName}_buid_docker_new.md5 /tmp/${ProjectName}_buid_docker_old.md5
