// author: JGZ
// time:   2021-01-18 15:16
package generate

type Data struct {
	Mode         *ModelData
	New          *NewData
	RepoDataList []*RepoData
}

type ModelData struct {
	PackageName string
	Imports     []string
	Entitys     []*ModelEntity
}

type ModelEntity struct {
	EntityName       string
	TableName        string
	EntityAnnotation string
	EntityFields     []*ModelEntityField
}

type ModelEntityField struct {
	FieldName       string
	FieldType       string
	ColumnName      string
	ColumnType      string
	FieldAnnotation string
}

type NewData struct {
	PackageName string
	Imports     []string
	Entitys     []*NewEntity
}

type NewEntity struct {
	EntityName     string
	PrimaryKeyType string
}

type NewEntityField struct {
}

type RepoData struct {
	PackageName          string
	Imports              []string
	EntityName           string
	TableName            string
	PrimaryKeyName       string
	PrimaryKeyType       string
	PrimaryKeyColumnName string
	PrimaryKeyColumnType string
	ColumnNames          []string
	EntityFieldMap       map[string]string
	EntityFields         []*RepoEntityField
}

type RepoEntityField struct {
	FieldName  string
	FieldType  string
	ColumnName string
	ColumnType string
}
