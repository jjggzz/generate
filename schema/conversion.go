// author: JGZ
// time:   2021-01-15 17:40
// 将数据库中的结构转换成为模板需要的实体结构
package schema

import "strings"

// 将从数据库中获取的结构转换为实体结构
func ConversionTableToEntity(tables []*Table) []*Entity {
	entitys := make([]*Entity, 0, len(tables))
	for _, e := range tables {
		entitys = append(entitys, conversionTable(e))
	}
	return entitys
}

// 表转换为实体
func conversionTable(table *Table) *Entity {
	entity := new(Entity)
	entity.EntityName = UnderlineToUpCamel(table.TableName)
	entity.TableName = table.TableName
	entity.EntityAnnotation = table.TableComment
	entityFields := make([]*EntityField, 0, len(table.TableFields))
	columnNames := make([]string, 0, len(table.TableFields))
	for _, ee := range table.TableFields {
		field := conversionField(ee)
		entityFields = append(entityFields, field)
		columnNames = append(columnNames, field.ColumnName)
		if field.ColumnKey == "PRI" {
			entity.PrimaryKeyColumnName = field.ColumnName
			entity.PrimaryKeyColumnType = field.ColumnType
			entity.PrimaryKeyName = field.FieldName
			entity.PrimaryKeyType = field.FieldType
		}
	}
	entity.ColumnNames = columnNames
	entity.EntityFields = entityFields
	return entity
}

// 表字段转换为实体字段
func conversionField(field *TableField) *EntityField {
	entityField := new(EntityField)
	entityField.FieldName = UnderlineToUpCamel(field.ColumnName)
	entityField.FieldType = conversionColumnTypeToEntityType(field.ColumnType, field.IsNullable == "YES")
	entityField.ColumnName = field.ColumnName
	entityField.ColumnType = field.ColumnType
	entityField.ColumnKey = field.ColumnKey
	entityField.FieldAnnotation = field.ColumnComment
	return entityField
}

// 转换为首字母大写
// str:		源字符串
// example:	abc => Abc
func StrFirstLetterToCapital(str string) string {
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
func StrFirstLetterToLowercase(str string) string {
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
func UnderlineToUpCamel(str string) string {
	split := strings.Split(str, "_")
	for index, subStr := range split {
		split[index] = StrFirstLetterToCapital(subStr)
	}
	return strings.Join(split, "")
}

func conversionColumnTypeToEntityType(columnType string, isNullable bool) string {
	split := strings.Split(columnType, "(")
	switch strings.ToLower(split[0]) {
	case "bit":
		return simpleIF(isNullable, "sql.NullBool", "bool").(string)
	case "tinyint":
		return simpleIF(isNullable, "sql.NullByte", "byte").(string)
	case "smallint":
		return simpleIF(isNullable, "sql.NullInt16", "int16").(string)
	case "int", "integer":
		return simpleIF(isNullable, "sql.NullInt32", "int32").(string)
	case "bigint":
		return simpleIF(isNullable, "sql.NullInt64", "int64").(string)
	case "float", "double", "decimal":
		return simpleIF(isNullable, "sql.NullFloat64", "float64").(string)
	case "char", "varchar", "tinyblob", "tinytext", "mediumblob", "mediumtext", "longblob", "longtext":
		return simpleIF(isNullable, "sql.NullString", "string").(string)
	case "date", "datetime", "timestamp":
		return simpleIF(isNullable, "sql.NullTime", "time.Time").(string)
	default:
		return "interface{}"
	}
}

func simpleIF(condition bool, value1 any, value2 any) any {
	if condition {
		return value1
	} else {
		return value2
	}
}
