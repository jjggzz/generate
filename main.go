package main

import (
	"bytes"
	"github.com/jjggzz/generate/build"
	"github.com/jjggzz/generate/temp"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	tables, err := build.Build("customer")
	if err != nil {
		panic(err)
	}
	entitys := build.ConversionTableToEntity(tables)

	tmp := template.New("demo")
	text, err := readFile("D:\\code\\GoCode\\generate\\temp\\temp.gotemplate")
	if err != nil {
		panic(err)
	}
	parse, err := tmp.Parse(text)
	if err != nil {
		panic(err)
	}
	file, err := os.Create("./demo/demo.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data := temp.NewData("demo", "xorm", entitys)
	buffer := bytes.NewBuffer(nil)
	err = parse.Execute(buffer, data)
	source, err := format.Source(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = file.Write(source)
	if err != nil {
		panic(err)
	}
}

func readFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
