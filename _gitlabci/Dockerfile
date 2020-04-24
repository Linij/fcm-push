FROM hub.591.com.tw/centos-base:7

ARG CI_COMMIT_REF_NAME
ARG ProjecHtdocsDir
RUN cd /tmp/ && wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz && \
    tar -zxvf go1.14.2.linux-amd64.tar.gz && \
    mv /tmp/go /usr/local/go && \
    ln -s  /usr/local/go/bin/* /usr/bin/
COPY ./_gitlabci/start.sh /start.sh
WORKDIR ${ProjecHtdocsDir}
EXPOSE 80 443 22


CMD ["/bin/bash","/start.sh"]
