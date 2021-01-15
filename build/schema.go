// author: JGZ
// time:   2021-01-15 17:40
// 加载数据库中的表及字段结构
package build

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
)

type Table struct {
	TableName    string
	TableComment string
	TableFields  []*TableField
}

type TableField struct {
	ColumnName    string
	ColumnType    string
	ColumnComment string
}

var db *sql.DB

func init() {
	conn, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/customer?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = conn
}

// 关闭资源
// closer: 目标资源
func CloseResource(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Println("close resource err: ", err)
	}
}

const getTableSql = "SELECT TB.TABLE_NAME,TB.TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES TB Where TB.TABLE_SCHEMA = ?"
const getFieldSql = "SELECT COL.COLUMN_NAME,COL.COLUMN_TYPE,COL.COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS COL Where COL.TABLE_NAME = ?"

// 加载表结构
// schema: 数据库名
func Build(schema string) ([]*Table, error) {
	tables, err := getTables(schema)
	if err != nil {
		return nil, err
	}
	for _, e := range tables {
		fields, err := getField(e.TableName)
		if err == nil {
			e.TableFields = fields
		}
	}
	return tables, err
}

func getTables(schema string) ([]*Table, error) {
	rows, err := db.Query(getTableSql, schema)
	if err != nil {
		panic(fmt.Sprintf("query tables err: %s", err.Error()))
	}
	defer CloseResource(rows)
	tables := make([]*Table, 0, 10)
	for rows.Next() {
		table := new(Table)
		err := rows.Scan(&table.TableName, &table.TableComment)
		if err != nil {
			log.Printf("query tables err: %s", err.Error())
		} else {
			tables = append(tables, table)
		}
	}
	return tables, nil
}

func getField(table string) ([]*TableField, error) {
	rows, err := db.Query(getFieldSql, table)
	if err != nil {
		panic(fmt.Sprintf("query field err: %s", err.Error()))
	}
	defer CloseResource(rows)
	fields := make([]*TableField, 0, 10)
	for rows.Next() {
		field := new(TableField)
		err := rows.Scan(&field.ColumnName, &field.ColumnType, &field.ColumnComment)
		if err != nil {
			log.Printf("query field err: %s", err.Error())
		} else {
			fields = append(fields, field)
		}
	}
	return fields, nil
}
