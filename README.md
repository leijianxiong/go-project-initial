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
wget https://golang.google.cn/dl/go1.14.6.linux-amd64.tar.gz
#目标机器
tar -C /usr/local -xzf go1.14.6.linux-amd64.tar.gz
```
- 设置go相关环境变量:
```
vim $HOME/.profile
#sudo vim /etc/profile.d/go.sh
#go env
export GO111MODULE=on 
export GOPROXY=https://goproxy.cn 
export GOROOT=/usr/local/go 
export GOPATH=$HOME/go
#export GOPATH=/data/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN

source $HOME/.profile
#sudo source /etc/profile.d/go.sh
``` 
- 检查go环境是否安装正常: `go env`
- 拉取代码: `git pull`
- 项目配置: `cd <project-name> && cp .env.example .env`
- 拉取包: `go mod tidy`
- 数据库迁移参考: [docs/migrate.md](docs/migrate.md)
```bash
新建数据库 dbname
#导入goadmin数据库文件
mysql -uroot -p <dbname> < $GOPATH/pkg/mod/github.com/\!go\!admin\!group/go-admin@v1.2.14/data/admin.sql

#数据库迁移
#下载包
go get github.com/golang-migrate/migrate/v4@v4.11.0
#进入包目录
cd $GOPATH/pkg/mod/github.com/golang-migrate/migrate/v4@v4.11.0/cmd/migrate/
#编译
go build -o $GOBIN/migrate -ldflags="-X main.Version=4.11.0" -tags "mysql file"
#迁移
#migrate -verbose -database "mysql://root@root123@tcp(192.168.33.22:3306)/staronline?charset=utf8mb4&parseTime=True&loc=Local" -path migrations up
migrate -verbose -database "mysql://username:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local" -path migrations up

```
- `nginx` 文件配置 [nginx-proxy.conf.example](nginx-proxy.conf.example)
```
cp nginx-proxy.conf.example nginx-proxy.conf #修改对应配置
sudo ln -s `pwd`/nginx-proxy.conf /etc/nginx/conf.d/<project-name>-nginx-proxy.conf
nginx -t
nginx -s reload
```
- 编译: `cd cmd/entry/ && go build -o cmd/entry/entry-<project-name>.cmd`
- 运行: `./entry-<project-name>.cmd &`

### 代码更新
- git pull 
- cd cmd/entry/ && go build -o entry-<project-name>.cmd
- ps -ef | grep entry-<project-name>.cmd 得到PID
- kill -s SIGHUP `<PID>`

### 配置文件改动
- ps -ef | grep entry-<project-name>.cmd 得到PID
- kill -s SIGHUP `<PID>`
