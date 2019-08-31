package builder

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
	"text/template"
)

const (
	StaticFilePath     = "static.go"
	StaticFileTemplate = `
package {{with .PackageName}}{{index .}}{{else}}builder{{end}}

var (
{{range $key, $value := .StaticVariables}}{{$key}} = &#96;{{$value}}&#96;{{end}}
)
`
)

var (
	DefaultFilePath = StaticFilePath
)

func Build(staticVariables map[string]interface{}) (result string, err error) {
	bytesBuffer := bytes.Buffer{}

	tmpl := template.Must(template.New(DefaultFilePath).Parse(StaticFileTemplate))
	err = tmpl.Execute(&bytesBuffer, staticVariables)
	result = bytesBuffer.String()
	result = strings.Replace(result, "`", "", -1)

	return html.UnescapeString(result), err
}
