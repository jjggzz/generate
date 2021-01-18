package main

import (
	"flag"
	"github.com/jjggzz/generate/generate"
	"github.com/jjggzz/generate/schema"
	"log"
	"path"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "db server ip")
	port        = flag.Int("port", 3306, "db server port")
	userName    = flag.String("u", "", "db server username")
	password    = flag.String("p", "", "db server password")
	schema      = flag.String("schema", "", "db server name")
	packageName = flag.String("pkg", "repository", "go package name")
	outputPath  = flag.String("out", "./", "file output path")
)

func main() {
	flag.Parse()
	if *userName == "" {
		log.Println("place input username")
		return
	}
	if *password == "" {
		log.Println("place input password")
		return
	}
	if *schema == "" {
		log.Println("place input name")
		return
	}

	schema.Init(*userName, *password, *ip, *port, *schema)
	tables, err := schema.Load(*schema)
	if err != nil {
		panic(err)
	}
	entitys := schema.ConversionTableToEntity(tables)
	data := generate.New(*packageName, entitys)
	generate.Generate(path.Join(*outputPath, *packageName), data)
}
