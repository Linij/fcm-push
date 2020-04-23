# FCM PUSH
Google FCM push use go &amp; iris framework .


# Currently under development, please do not use


# 环境要求
awk

go 1.14+ 需要配置相关环境变量

# 启动命令，进入项目根目录
export GO111MODULE=on

nohup go run main.go 2>&1 >run.out &

# 需要域名
localhost:8089 -> push.debug.8591.com.hk

# 代码变动时需要重新编译
lsof -i:8089 | awk 'NR==2 {print $2}' | xargs kill

nohup go run main.go 2>&1 >run.out &
