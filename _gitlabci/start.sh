#! /bin/bash
ProjecHtdocsDir=/home/htdocs/gopush
SUPER_CONFIG_PATH=/super_config

export GO111MODULE=on
\cp -rf ${SUPER_CONFIG_PATH}/* /etc/supervisord.d/
/sbin/sshd
/usr/bin/supervisord
/usr/sbin/crond -n -p