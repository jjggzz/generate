// author: JGZ
// time:   2021-01-15 19:37
package generate

import "github.com/jjggzz/generate/schema"

func New(packageName string, entitys []*schema.Entity) *Data {
	data := &Data{}
	data.RepoDataList = buildRepoDataList(packageName, entitys)
	return data
}

func buildRepoDataList(packageName string, entitys []*schema.Entity) []*RepoData {
	data := make([]*RepoData, 0, len(entitys))

	for _, e := range entitys {
		repoData := new(RepoData)
		repoData.PackageName = packageName
		repoData.Imports = buildRepoImports(e)
		repoData.PrimaryKeyName = e.PrimaryKeyName
		repoData.PrimaryKeyType = e.PrimaryKeyType
		repoData.PrimaryKeyColumnName = e.PrimaryKeyColumnName
		repoData.PrimaryKeyColumnType = e.PrimaryKeyColumnType
		repoData.EntityName = e.EntityName
		repoData.TableName = e.TableName
		repoData.EntityAnnotation = e.EntityAnnotation
		repoData.ColumnNames = e.ColumnNames
		fields := make([]*RepoEntityField, 0, len(e.EntityFields))
		strings := make(map[string]string)
		for _, ee := range e.EntityFields {
			field := new(RepoEntityField)
			field.FieldName = ee.FieldName
			field.ColumnName = ee.ColumnName
			field.ColumnType = ee.ColumnType
			field.FieldType = ee.FieldType
			field.FieldAnnotation = ee.FieldAnnotation
			fields = append(fields, field)
			strings[ee.ColumnName] = ee.FieldName
		}
		repoData.EntityFieldMap = strings
		repoData.EntityFields = fields
		data = append(data, repoData)
	}
	return data
}

// 构建repo模板文件的数据
func buildRepoImports(entitys *schema.Entity) []string {
	imports := make([]string, 0)
	imports = append(imports, "_ \"github.com/go-sql-driver/mysql\"")
	imports = append(imports, "\"github.com/jmoiron/sqlx\"")
	imports = append(imports, "\"database/sql\"")
	imports = append(imports, "\"strings\"")
	for _, e := range entitys.EntityFields {
		switch e.FieldType {
		case "*time.Time":
			if !inStringSilce(imports, "\"time\"") {
				imports = append(imports, "\"time\"")
			}
		default:
			continue
		}
	}
	return imports
}

func inStringSilce(silce []string, str string) bool {
	for _, v := range silce {
		if v == str {
			return true
		}
	}
	return false
}
