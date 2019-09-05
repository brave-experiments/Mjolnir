package terra

import "github.com/tidwall/gjson"

const (
	ModulesLocator = "modules.outputs"
)

type OutputRecord struct {
	Name      string
	Sensitive bool        `json:"sensitive"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
}

func (outputRecord *OutputRecord) ParseOutputsFromJson(jsonBody string) {
	outputLocator := ModulesLocator + "." + outputRecord.Name
	jsonOutputs := gjson.Get(jsonBody, outputLocator)

	if false == jsonOutputs.Exists() {
		return
	}

	//outputRecord.Sensitive = jsonOutputs.Exists()
}
