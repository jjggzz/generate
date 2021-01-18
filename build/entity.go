// author: JGZ
// time:   2021-01-15 18:21
// 实体结构
package build

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
	OrmTag          string
	JsonTag         string
	FieldAnnotation string
}
