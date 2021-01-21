# generate
golang mysql操作代码生成工具，采用sqlx作为操作库
```
# mysql驱动安装

# sqlx安装
go get github.com/jmoiron/sqlx
```
# windows安装
## 1.执行wininstall.bat
```
  > wininstall.bat
```
会在%GOPATH%\go\bin目录下生成sqlxgen.exe
## 2.添加环境变量
将%GOPATH%\go\bin添加到环境变量

```
  # 查看帮助
  >sqlxgen.exe --help
  Usage of sqlxgen.exe:
  -db string
        db server name
  -ip string
        db server ip (default "127.0.0.1")
  -out string
        file output path (default "./")
  -p string
        db server password
  -pkg string
        go package name (default "repository")
  -port int
        db server port (default 3306)
  -u string
        db server username
```
# linux安装
## 1.执行linuxinstall.sh
```
  > ./linuxinstall.sh
```
会在%GOPATH%/go/bin目录下生成sqlxgen
## 2.添加环境变量
将%GOPATH%/go/bin添加到环境变量
```
  # 查看帮助
  >sqlxgen --help
  Usage of sqlxgen:
  -db string
        db server name
  -ip string
        db server ip (default "127.0.0.1")
  -out string
        file output path (default "./")
  -p string
        db server password
  -pkg string
        go package name (default "repository")
  -port int
        db server port (default 3306)
  -u string
        db server username
```
# 使用
在你的项目路径下执行
```
# linux
> sqlxgen -u root -p 123456 -db dbname
```
```
# windows
> sqlxgen.exe -u root -p 123456 -db dbname
```



