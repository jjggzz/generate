package main

import (
	"flag"
	"log"
	"path"

	"github.com/jjggzz/generate/generate"
	"github.com/jjggzz/generate/schema"
)

var (
	ip          = flag.String("ip", "127.0.0.1", "db server ip")
	port        = flag.Int("port", 3306, "db server port")
	userName    = flag.String("u", "", "db server username")
	password    = flag.String("p", "", "db server password")
	db          = flag.String("db", "", "db server name")
	packageName = flag.String("pkg", "repository", "go package name")
	outputPath  = flag.String("out", "./", "file output path")
	table       = flag.String("table", "", "if not specified generate all")
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
	if *db == "" {
		log.Println("place input name")
		return
	}

	schema.Init(*userName, *password, *ip, *port, *db)
	// schema.Init("root", "iSm!23465", "mysql", 3306, "chengdu")
	tables, err := schema.Load(*db, *table)
	// tables, err := schema.Load("demo", "")
	if err != nil {
		panic(err)
	}
	entitys := schema.ConversionTableToEntity(tables)
	data := generate.New(*packageName, entitys)
	generate.Generate(path.Join(*outputPath, *packageName), data)
}
