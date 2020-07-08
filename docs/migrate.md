```
#进入包目录
cd $GOPATH/pkg/mod/github.com/golang-migrate/migrate/v4@v4.11.0/cmd/migrate/
#编译
go build -o $GOBIN/migrate -ldflags="-X main.Version=4.11.0" -tags "mysql file"

#创建文件
migrate create -ext sql -dir migrations <name>

#迁移
migrate -verbose -database "mysql://root:123456@tcp(192.168.33.22:3306)/wallet_services?charset=utf8mb4&parseTime=True&loc=Local" -path migrations up

#遇到错误时会记录具体的v是dirty的使用 force V 忽略错误, 后续再次执行migrate:  force V     Set version V but don't run migration (ignores dirty state)
migrate -verbose -database "mysql://root:123456@tcp(192.168.33.22:3306)/wallet_services?charset=utf8mb4&parseTime=True&loc=Local" -path migrations force <V>

#回退
migrate -verbose -database "mysql://root:123456@tcp(192.168.33.22:3306)/wallet_services?charset=utf8mb4&parseTime=True&loc=Local" -path migrations down 1

#回退到这一个版本(已应用该版本)
migrate -verbose -database "mysql://root:123456@tcp(192.168.33.22:3306)/wallet_services?charset=utf8mb4&parseTime=True&loc=Local" -path migrations goto 20200621195734
```
