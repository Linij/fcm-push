FROM hub.591.com.tw/centos-go:7

ARG CI_COMMIT_REF_NAME
ARG ProjecHtdocsDir
COPY ./_gitlabci/start.sh /start.sh
WORKDIR ${ProjecHtdocsDir}
EXPOSE 80 443 22


CMD ["/bin/bash","/start.sh"]
