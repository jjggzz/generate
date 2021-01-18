// author: JGZ
// time:   2021-01-15 18:21
// 实体结构
package schema

type Entity struct {
	EntityName       string
	TableName        string
	PrimaryKeyName   string
	PrimaryKeyType   string
	EntityAnnotation string
	ColumnNames      []string
	EntityFields     []*EntityField
}

type EntityField struct {
	FieldName       string
	FieldType       string
	ColumnName      string
	ColumnType      string
	ColumnKey       string
	FieldAnnotation string
}
