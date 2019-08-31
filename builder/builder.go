package builder

import (
	"bytes"
	"fmt"
	"html/template"
)

const (
	StaticFilePath     = "static.go"
	StaticFileTemplate = `
package %s

var (
{{range $key, $value := .StaticVariables}}{{$key}} = {{$value}}{{end}}
)
`
)

var (
	DefaultFilePath = StaticFilePath
)

func Build(packageName string, staticVariables map[string]interface{}) (result string, err error) {
	staticFileTemplate := fmt.Sprintf(StaticFileTemplate, packageName)
	bytesBuffer := bytes.Buffer{}
	tmpl := template.Must(template.New(DefaultFilePath).Parse(staticFileTemplate))
	err = tmpl.Execute(&bytesBuffer, staticVariables)
	result = bytesBuffer.String()

	return result, err
}
