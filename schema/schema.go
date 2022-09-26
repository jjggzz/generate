// author: JGZ
// time:   2021-01-15 17:40
// 加载数据库中的表及字段结构
package schema

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Table struct {
	TableName    string
	TableComment string
	TableFields  []*TableField
}

type TableField struct {
	ColumnName    string
	ColumnKey     string
	ColumnType    string
	IsNullable    string
	ColumnComment string
}

type Entity struct {
	EntityName           string
	TableName            string
	PrimaryKeyName       string
	PrimaryKeyType       string
	PrimaryKeyColumnName string
	PrimaryKeyColumnType string
	EntityAnnotation     string
	ColumnNames          []string
	EntityFields         []*EntityField
}

type EntityField struct {
	FieldName       string
	FieldType       string
	ColumnName      string
	ColumnType      string
	ColumnKey       string
	FieldAnnotation string
}

var db *sql.DB

func Init(userName string, password string, ip string, port int, schema string) {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		userName,
		password,
		ip,
		port,
		schema)
	conn, err := sql.Open("mysql", url)
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
const getFieldSql = "SELECT COL.COLUMN_NAME,COL.COLUMN_TYPE,COL.COLUMN_KEY,COL.IS_NULLABLE,COL.COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS COL Where COL.TABLE_SCHEMA = ? and COL.TABLE_NAME = ?"

// 加载表结构
// schema: 数据库名
func Load(schema string, tableName string) ([]*Table, error) {
	var tables []*Table
	var err error
	if tableName == "" {
		tables, err = getTables(schema)
	} else {
		tableNameList := strings.Split(tableName, ",")
		tables, err = getTablesFilterTableName(schema, tableNameList)
	}
	if err != nil {
		return nil, err
	}
	for _, e := range tables {
		fields, err := getField(schema, e.TableName)
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

func getTablesFilterTableName(schema string, tableName []string) ([]*Table, error) {
	rows, err := db.Query(getTableSql, schema)
	if err != nil {
		panic(fmt.Sprintf("query tables err: %s", err.Error()))
	}
	defer CloseResource(rows)

	fm := make(map[string]int)
	for i, v := range tableName {
		fm[v] = i
	}
	tables := make([]*Table, 0, 10)
	for rows.Next() {
		table := new(Table)
		err := rows.Scan(&table.TableName, &table.TableComment)
		if err != nil {
			log.Printf("query tables err: %s", err.Error())
		} else {
			if _, ok := fm[table.TableName]; ok {
				tables = append(tables, table)
			}
		}
	}
	return tables, nil
}

func getField(schema string, table string) ([]*TableField, error) {
	rows, err := db.Query(getFieldSql, schema, table)
	if err != nil {
		panic(fmt.Sprintf("query field err: %s", err.Error()))
	}
	defer CloseResource(rows)
	fields := make([]*TableField, 0, 10)
	for rows.Next() {
		field := new(TableField)
		rows.Scan()
		err := rows.Scan(&field.ColumnName, &field.ColumnType, &field.ColumnKey, &field.IsNullable, &field.ColumnComment)
		if err != nil {
			log.Printf("query field err: %s", err.Error())
		} else {
			fields = append(fields, field)
		}
	}
	return fields, nil
}
