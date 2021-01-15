// author: JGZ
// time:   2021-01-15 18:21
// 实体结构
package build

type Entity struct {
	EntityName       string
	EntityAnnotation string
	EntityFields     []*EntityField
}

type EntityField struct {
	FieldName       string
	FieldType       string
	FieldAnnotation string
}