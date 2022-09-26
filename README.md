# generate
golang mysql数据操作层代码生成工具，采用sqlx作为操作库，在项目中想使用该工具需要为你的项目安装sqlx依赖：
```
# mysql驱动安装
go get -u github.com/go-sql-driver/mysql
# sqlx安装
go get -u github.com/jmoiron/sqlx
```
生成的代码结构如下
```
  repository
    -- demo_repository.go # 仓库
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
  -table string
        if not specified generate all
```
# linux安装
## 1.执行linuxinstall.sh
```
  > chmod 777 linuxinstall.sh
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
  -table string
        if not specified generate all
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



