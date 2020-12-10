## Model
把数据库的表转换成`go`语言的`struct`，支持 `gorm`    

### 安装 
安装到`GOPATH`的 `bin`目录.
```
go get -u github.com/wg2019/model
```
#### 异常
1. 拉不到依赖
    ```
    go: downloading github.com/wg2019/model v0.0.0-20201210031901-896455bf9526
    go: github.com/wg2019/model upgrade => v0.0.0-20201210031901-896455bf9526
    go get: github.com/wg2019/model@v0.0.0-20201210031901-896455bf9526 requires
        github.com/jinzhu/gorm@v1.9.16 requires
        golang.org/x/crypto@v0.0.0-20191205180655-e7c4368fe9dd: unrecognized import path "golang.org/x/crypto": https fetch: Get "https://golang.org/x/crypto?go-get=1": dial tcp 216.239.37.1:443: i/o timeout
    ```
    解决：
    ```
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
    ```
### 帮助
```
NAME:
   model - 实体生成器

USAGE:
   model [global options] [arguments...]

DESCRIPTION:
   啊吧啊吧啊吧

GLOBAL OPTIONS:
   --host value, -h value        主机 (default: "127.0.0.1")
   --port value, -P value        端口 (default: 3306)
   --output value, -o value      生成文件夹 (default: "./")
   --package value, --pkg value  包名称 (default: "db")
   --database value, --db value  数据库名称
   --table value, -t value       表名称
   --user value, -u value        用户名称 (default: "root")
   --password value, -p value    密码
   --desc value                  描述 (default: "json")
```
#### example:
```
go run . -h localhost -P 3306 -o ./ -db database_test -u root -p pwd123
```
