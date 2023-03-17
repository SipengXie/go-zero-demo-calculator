# gonivinck
一个基于 `docker` 的 `go-zero` 本地开发运行环境。


## 使用
### 1. 按需修改 .env 配置
~~~env
# 设置时区
TZ=Asia/Shanghai
# 设置网络模式
NETWORKS_DRIVER=bridge

# PATHS ##########################################
# 宿主机上代码存放的目录路径
CODE_PATH_HOST=./code

# ETCD ###########################################
# Etcd 服务映射宿主机端口号，可在宿主机127.0.0.1:2379访问
ETCD_PORT=2379

ETCD_MANAGE_PORT=7000
~~~

### 2.启动服务
- 启动全部服务
```bash
docker-compose up -d
```
- 按需启动部分服务
```bash
docker-compose up -d etcd golang mysql redis
```

### 3.运行代码
将开发代码放置 `CODE_PATH_HOST` 指定的本机目录，进入 `golang` 容器，运行项目代码。
~~~bash
docker exec -it gonivinck_golang_1 bash
~~~

### 4.代码简介
- 放置在 `code` 文件夹下是一个Demo代码，它分为 `adder` 和 `calc` 两个部分。其中 `adder` 是一个 rpc 服务，而 `calc` 则向外提供 api 接口。
- 具体教程可以参考：
    - https://juejin.cn/post/7036011047391592485 （注意看评论区的讨论）

    - https://go-zero.dev/cn/docs/quick-start/micro-service （以及后面的进阶教程）


 `calc` 和 `adder` 的连接通过 `etcd` 实现，这一点可以从双方文件夹下 `etc/xxx.yaml` 看出来。
 
 一般来说，在配置完 `etc/xxx.yaml` 文件后会去更新 `internal/config/config.go` 以及 `internal/svc/servicecontext.go` 来获得 rpc 连接。