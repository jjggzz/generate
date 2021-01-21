@ECHO ON
go get -u github.com/go-sql-driver/mysql@v1.5.0
go get -u github.com/kevinburke/go-bindata/go-bindata@v3.22.0+incompatible
go install github.com/jjggzz/generate/cmd/sqlxgen
@ECHO OFF

