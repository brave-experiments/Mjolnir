package main

import (
	"bytes"
	"fmt"
	"github.com/brave-experiments/apollo-devops/terra"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	StaticKeyPrefix    = "Static"
	StaticFilePath     = "terra/static.go"
	StaticFileTemplate = `
package {{with .PackageName}}{{index .}}{{else}}builder{{end}}

var (
{{range $key, $value := .StaticVariables}}Static{{$key | Title}} = &#96;{{$value}}&#96;{{end}}
)
`
)

var (
	DefaultFilePath = StaticFilePath
)

func Build(staticVariables map[string]interface{}) (result string, err error) {
	funcMap := template.FuncMap{
		"Title": strings.Title,
	}

	bytesBuffer := bytes.Buffer{}

	tmpl := template.Must(template.New(DefaultFilePath).Funcs(funcMap).Parse(StaticFileTemplate))
	err = tmpl.Execute(&bytesBuffer, staticVariables)
	result = bytesBuffer.String()
	result = strings.Replace(result, "`", "", -1)

	return html.UnescapeString(result), err
}

func main() {
	recipes := terra.DefaultRecipes
	staticVariables := make(map[string]string)

	for key, recipe := range recipes {
		err := recipe.ParseBody()

		if nil != err {
			log.Panicln(err)
		}

		staticVariables[key] = recipe.Body
	}

	staticVariablesMap := map[string]interface{}{
		"PackageName":     "terra",
		"StaticVariables": staticVariables,
	}

	result, err := Build(staticVariablesMap)

	if nil != err {
		log.Panicln(err)
	}

	file, err := os.Create(DefaultFilePath)

	if nil != err {
		log.Panicln(err)
	}

	bytesCount, err := file.WriteString(result)

	if nil != err {
		log.Panicln(err)
	}

	fmt.Printf("\nSuccessfully wrote %v bytes to path %s \n", bytesCount, DefaultFilePath)
}
