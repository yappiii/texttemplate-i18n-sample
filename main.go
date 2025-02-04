package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"os"

	"github.com/BurntSushi/toml"
)

//go:embed templates/*tmpl
var templStatic embed.FS

type TextToml struct {
	Hello string `toml:"hello"`
}

func main() {
	lang := os.Getenv("LANG")
	textDir := "templates/text/" + lang + ".toml"
	var config TextToml
	_, err := toml.DecodeFile(textDir, &config)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("main").ParseFS(templStatic, "templates/*tmpl")
	if err != nil {
		panic(err)
	}
	tmplConf := map[string]string{
		"Hello": config.Hello,
	}
	textBuf := &bytes.Buffer{}
	err = tmpl.ExecuteTemplate(textBuf, "main", tmplConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(textBuf.String())
}
