// author: JGZ
// time:   2021-01-15 19:37
package generate

import "github.com/jjggzz/generate/schema"

func New(packageName string, entitys []*schema.Entity) *Data {
	data := &Data{}
	data.Mode = &ModelData{
		PackageName: packageName,
		Imports:     buildModelImports(entitys),
		Entitys:     buildModelEntitys(entitys),
	}
	data.New = &NewData{
		PackageName: packageName,
		Imports:     buildNewImports(entitys),
		Entitys:     buildNewEntitys(entitys),
	}
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
		repoData.ColumnNames = e.ColumnNames
		fields := make([]*RepoEntityField, 0, len(e.EntityFields))
		strings := make(map[string]string)
		for _, ee := range e.EntityFields {
			field := new(RepoEntityField)
			field.FieldName = ee.FieldName
			field.ColumnName = ee.ColumnName
			field.ColumnType = ee.ColumnType
			field.FieldType = ee.FieldType
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
	for _, e := range entitys.EntityFields {
		switch e.FieldType {
		case "time.Time":
			imports = append(imports, "time")
		default:
			continue
		}
	}
	return imports
}

// 构建model模板文件的数据
func buildModelImports(entitys []*schema.Entity) []string {
	imports := make([]string, 0)
	for _, e := range entitys {
		for _, ee := range e.EntityFields {
			switch ee.FieldType {
			case "time.Time":
				imports = append(imports, "time")
			default:
				continue
			}
		}
	}
	return imports
}
func buildModelEntitys(entitys []*schema.Entity) []*ModelEntity {
	entities := make([]*ModelEntity, 0, len(entitys))
	for _, e := range entitys {
		entity := new(ModelEntity)
		entity.TableName = e.TableName
		entity.EntityAnnotation = e.EntityAnnotation
		entity.EntityName = e.EntityName
		fields := make([]*ModelEntityField, 0, len(e.EntityFields))
		for _, ee := range e.EntityFields {
			field := new(ModelEntityField)
			field.ColumnName = ee.ColumnName
			field.FieldName = ee.FieldName
			field.ColumnType = ee.ColumnType
			field.FieldType = ee.FieldType
			field.FieldAnnotation = ee.FieldAnnotation
			fields = append(fields, field)
		}
		entity.EntityFields = fields
		entities = append(entities, entity)
	}
	return entities
}

// 构建new模板文件的数据
func buildNewImports(entitys []*schema.Entity) []string {
	return make([]string, 0, 0)
}
func buildNewEntitys(entitys []*schema.Entity) []*NewEntity {
	entities := make([]*NewEntity, 0, len(entitys))
	for _, e := range entitys {
		entity := new(NewEntity)
		entity.EntityName = e.EntityName
		entity.PrimaryKeyType = e.PrimaryKeyType
		entities = append(entities, entity)
	}
	return entities
}
