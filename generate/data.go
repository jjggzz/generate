// author: JGZ
// time:   2021-01-18 15:16
package generate

type Data struct {
	RepoDataList []*RepoData
}

type RepoData struct {
	PackageName          string
	Imports              []string
	EntityName           string
	TableName            string
	EntityAnnotation     string
	PrimaryKeyName       string
	PrimaryKeyType       string
	PrimaryKeyColumnName string
	PrimaryKeyColumnType string
	ColumnNames          []string
	EntityFieldMap       map[string]string
	EntityFields         []*RepoEntityField
}

type RepoEntityField struct {
	FieldName       string
	FieldType       string
	ColumnName      string
	ColumnType      string
	FieldAnnotation string
}
