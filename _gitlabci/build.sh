#!/bin/bash
if [ "$CI_COMMIT_REF_NAME" == "online" ]
then
    CI_COMMIT_REF_NAME="online"
else
    CI_COMMIT_REF_NAME="develop"
fi
ProjectDir="${HtdocsDir}/${Project}-${CI_COMMIT_REF_NAME}"
function CheckoutREf()
{
    echo [准备切换到$CI_COMMIT_REF_NAME分支]
    if [ "$CI_COMMIT_REF_NAME" == "online" ]
    then
        git checkout $CI_COMMIT_REF_NAME
    else
        git checkout origin/$CI_COMMIT_REF_NAME -b $CI_COMMIT_REF_NAME
    fi
}

function UpdateConfig()
{
    if [ ! -d "update-src/${CI_COMMIT_REF_NAME}" ]
    then
        echo [未发现update-src/${CI_COMMIT_REF_NAME}文件夹]
    else
        echo [准备同步update-src/${CI_COMMIT_REF_NAME}文件夹]
        rsync -avp update-src/${CI_COMMIT_REF_NAME}/. ./
    fi
}

function UpdateEnv()
{
    if [ ! -f "config/config.toml.${CI_COMMIT_REF_NAME}.example" ]
    then
        echo [未发现config.toml.${CI_COMMIT_REF_NAME}.example文件]
    else
        echo [准备同步config.toml.${CI_COMMIT_REF_NAME}.example文件]
        rsync -avp config/config.toml.${CI_COMMIT_REF_NAME}.example config/config.toml
    fi
}

if [ ! -d ${ProjectDir} ]
then
    echo [准备创建文件夹${ProjectDir}]
    cp -ap $CI_PROJECT_DIR ${ProjectDir}
    cd ${ProjectDir}
    git config --global user.name "git-runner"
    echo [修改仓库拉取用户]
    sed 's|gitlab-ci-token:.*@|www:xnkyfekgy4TAgtV6kpYz@|g' -i ./.git/config
    CheckoutREf
    echo [准备获取最新的版本]
    git remote prune origin
    git fetch --all
    git reset --hard origin/$CI_COMMIT_REF_NAME
    git pull
    UpdateConfig
    UpdateEnv
else
    cd ${ProjectDir}
    git config --global user.name "git-runner"
    if [ -z "`git branch --list $CI_COMMIT_REF_NAME|grep \*`" ]
    then
        echo [检测到当前分支不是"$CI_COMMIT_REF_NAME",是"`git branch --list|grep \*`"]
        echo [准备获取最新的版本]
        git pull
    else
        echo [检测到当前分支是"$CI_COMMIT_REF_NAME"]
        echo [准备获取最新的版本]
        git remote prune origin
        git fetch --all
        git reset --hard origin/$CI_COMMIT_REF_NAME
        git pull
    fi
    UpdateConfig
    UpdateEnv
fi
