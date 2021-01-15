// author: JGZ
// time:   2021-01-15 17:40
// 将数据库中的结构转换成为模板需要的实体结构
package build

import "strings"

// 将从数据库中获取的结构转换为实体结构
func ConversionTableToEntity(tables []*Table) []*Entity {
	entitys := make([]*Entity, 0, len(tables))
	for _, e := range tables {
		entity := new(Entity)
		entity.EntityName = underlineToUpCamel(e.TableName)
		entity.EntityAnnotation = e.TableComment
		entityFields := make([]*EntityField, 0, len(e.TableFields))
		for _, ee := range e.TableFields {
			entityField := new(EntityField)
			entityField.FieldName = underlineToUpCamel(ee.ColumnName)
			entityField.FieldAnnotation = ee.ColumnComment
			entityField.FieldType = conversionColumnTypeToEntityType(ee.ColumnType)
			entityFields = append(entityFields, entityField)
		}
		entity.EntityFields = entityFields
		entitys = append(entitys, entity)
	}
	return entitys
}

// 转换为首字母大写
// str:		源字符串
// example:	abc => Abc
func strFirstLetterToCapital(str string) string {
	runes := []rune(str)
	for i, e := range runes {
		if i == 0 {
			if e >= 97 && e <= 122 {
				runes[i] = e - 32
			}
		}
	}
	return string(runes)
}

// 转换为首字母小写
// str:		源字符串
// example:	abc => abc
// example:	Abc => abc
func strFirstLetterToLowercase(str string) string {
	runes := []rune(str)
	for i, e := range runes {
		if i == 0 {
			if e >= 65 && e <= 91 {
				runes[i] = e + 32
			}
		}
	}
	return string(runes)
}

// 将下划线命名法字符串转换为大驼峰
// str:		源字符串
// example:	hello_world => HelloWorld
func underlineToUpCamel(str string) string {
	split := strings.Split(str, "_")
	for index, subStr := range split {
		split[index] = strFirstLetterToCapital(subStr)
	}
	return strings.Join(split, "")
}

func conversionColumnTypeToEntityType(columnType string) string {
	split := strings.Split(columnType, "(")
	switch strings.ToLower(split[0]) {
	case "bit":
		return "bool"
	case "tinyint":
		return "int8"
	case "smallint":
		return "int16"
	case "int", "integer":
		return "int32"
	case "bigint":
		return "int64"
	case "float":
		return "float32"
	case "double", "decimal":
		return "float64"
	case "char", "varchar", "tinyblob", "tinytext", "mediumblob", "mediumtext", "longblob", "longtext":
		return "string"
	case "date", "datetime":
		return "time.Time"
	default:
		return "interface{}"
	}
}
