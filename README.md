# FCM PUSH
Google FCM push use go &amp; iris framework .


# Currently under development, please do not use


# 环境要求
##### awk

##### go 1.14+ 需要配置相关环境变量

# 首次安装（必须）
##### 1. 需要将 config 目录下面对应的配置文件复制一份为 config.toml
##### 2. 需要将 firebase 的 json 存一份到 config/serviceAccountKey.json `可以查阅 example 看怎么获取 json`
##### 3. cp config/.AutoMigrate.go config/AutoMigrate.go && go run AutoMigrate.go `将创建好数据表，在此之前需要配置好数据库相关的配置`

# 持续集成：进入项目根目录
##### export GO111MODULE=on

##### cd config && rm config.toml && cp config.toml.(dev/develop/online).example config.toml
 
##### go run main.go (super 管理)

# 需要域名
0.0.0.0:8089 -> push.debug.8591.com.hk

# 代码变动时需要重新编译
reload go run main.go

