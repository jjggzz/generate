// author: JGZ
// time:   2021-01-15 19:37
package temp

import "github.com/jjggzz/generate/build"

type Data struct {
	EntityPackageName string
	FrameworkName     string
	EntityImport      []string
	Entitys           []*build.Entity
}

func NewData(packageName string, frameworkName string, entitys []*build.Entity) *Data {

	return &Data{
		EntityPackageName: packageName,
		FrameworkName:     frameworkName,
		EntityImport:      buildEntityImport(entitys),
		Entitys:           entitys,
	}
}

// 构建实体导入的包
func buildEntityImport(entitys []*build.Entity) []string {
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
