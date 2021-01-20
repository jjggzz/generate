// author: JGZ
// time:   2021-01-16 14:56
package generate

import (
	"bytes"
	"fmt"
	"github.com/jjggzz/generate/schema"
	"github.com/jjggzz/generate/temp"
	"go/format"
	"os"
	"path"
	"strings"
	"text/template"
)

func Generate(basePath string, data *Data) {
	if !isDir(basePath) {
		err := os.Mkdir(basePath, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("create dir err: %s", err.Error()))
		}
	}

	maps := template.FuncMap{
		"StrFirstLetterToLowercase": schema.StrFirstLetterToLowercase,
		"join":                      strings.Join,
		"buildValueHold":            BuildValueHold,
		"findMap":                   func(key string, maps map[string]string) string { return maps[key] },
	}
	names := temp.AssetNames()
	for _, name := range names {

		switch name {
		case "model.gotemplate":
			tempText := temp.MustAssetString(name)
			tmp := template.New(name).Funcs(maps)
			parse, err := tmp.Parse(tempText)
			if err != nil {
				panic(fmt.Sprintf("template parse err: %s", err.Error()))
			}
			buffer := bytes.NewBuffer(nil)
			err = parse.Execute(buffer, data.Mode)
			source, err := format.Source(buffer.Bytes())
			split := strings.Split(name, ".")
			err = writeFile(path.Join(basePath, split[0]+".go"), source)
			if err != nil {
				panic(fmt.Sprintf("write file err: %s", err.Error()))
			}
		case "new.gotemplate":
			tempText := temp.MustAssetString(name)
			tmp := template.New(name).Funcs(maps)
			parse, err := tmp.Parse(tempText)
			if err != nil {
				panic(fmt.Sprintf("template parse err: %s", err.Error()))
			}
			buffer := bytes.NewBuffer(nil)
			err = parse.Execute(buffer, data.New)
			source, err := format.Source(buffer.Bytes())
			split := strings.Split(name, ".")
			err = writeFile(path.Join(basePath, split[0]+".go"), source)
			if err != nil {
				panic(fmt.Sprintf("write file err: %s", err.Error()))
			}
		case "repo.gotemplate":
			for _, ee := range data.RepoDataList {
				tempText := temp.MustAssetString(name)
				tmp := template.New(name).Funcs(maps)
				parse, err := tmp.Parse(tempText)
				if err != nil {
					panic(fmt.Sprintf("template parse err: %s", err.Error()))
				}
				buffer := bytes.NewBuffer(nil)
				err = parse.Execute(buffer, ee)
				source, err := format.Source(buffer.Bytes())
				split := strings.Split(name, ".")
				err = writeFile(path.Join(basePath, schema.StrFirstLetterToLowercase(ee.EntityName)+"_"+split[0]+".go"), source)
				if err != nil {
					panic(fmt.Sprintf("write file err: %s", err.Error()))
				}
			}
		case "example.gotemplate":
			for _, ee := range data.RepoDataList {
				tempText := temp.MustAssetString(name)
				tmp := template.New(name).Funcs(maps)
				parse, err := tmp.Parse(tempText)
				if err != nil {
					panic(fmt.Sprintf("template parse err: %s", err.Error()))
				}
				buffer := bytes.NewBuffer(nil)
				err = parse.Execute(buffer, ee)
				source, err := format.Source(buffer.Bytes())
				split := strings.Split(name, ".")
				err = writeFile(path.Join(basePath, schema.StrFirstLetterToLowercase(ee.EntityName)+"_"+split[0]+".go"), source)
				if err != nil {
					panic(fmt.Sprintf("write file err: %s", err.Error()))
				}
			}
		}

	}
}

func BuildValueHold(columns []string) string {
	str := ""
	for i := range columns {
		if i == 0 {
			str += "?"
		} else {
			str += ",?"
		}
	}
	return str
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func writeFile(path string, source []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer schema.CloseResource(file)
	_, err = file.Write(source)
	if err != nil {
		return err
	}
	return nil
}
