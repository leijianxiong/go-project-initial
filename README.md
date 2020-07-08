# go项目初始化

### 项目(project-name) 本地初始化步骤

调整每个项目模块名

```
git clone <url> <project-name>
cd <project-name>
vim go.mod
replace go-project-initial <project-name>
go mod tidy
github/gitlab/gitee create empty repository
git remote rename origin old-origin
git remote add origin <url>
git push -u origin
```

### 部署
- 安装go: doc golang download and [install](https://golang.google.cn/doc/install)
```
#上传/下载包
scp ~/Downloads/go1.14.4.linux-amd64.tar.gz eth@192.168.1.123:~/
#目标机器
tar -C /usr/local -xzf /home/eth/go1.14.4.linux-amd64.tar.gz
```
- 设置go相关环境变量:
```
vim $HOME/.profile
#go env
export GO111MODULE=on 
export GOPROXY=https://goproxy.cn 
export GOROOT=/usr/local/go 
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN

. $HOME/.profile
``` 
- 检查go环境是否安装正常: `go env`
- 拉取代码: `git pull`
- 项目配置: `cd <project-name> && cp .env.example .env`
- 拉取包: `go mod tidy`
- 执行数据库迁移参考: [docs/migrate.md](docs/migrate.md)
- ~~nginx配置文件目录加多一个软连接到 netquerysystem.conf, 修改配置~~ `cp nginx-proxy.conf.example nginx-proxy.conf`  [nginx-proxy.conf.example](nginx-proxy.conf.example)
- ~~nginx -s reload~~
- 编译: `go build -o cmd/entry/entry cmd/entry/main.go`
- 运行: `./cmd/entry/entry &`

### 代码更新
- git pull 
- go build -o cmd/entry/entry cmd/entry/main.go
- ps -ef | grep cmd/entry/entry 得到PID
- kill -s SIGHUP `<PID>`

### env改动
- ps -ef | grep cmd/entry/entry 得到PID
- kill -s SIGHUP `<PID>`
