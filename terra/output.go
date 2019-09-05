package terra

import (
	"github.com/tidwall/gjson"
)

const (
	ModulesLocator = "modules"
)

type OutputRecord struct {
	Name      string
	Sensitive bool   `json:"sensitive"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}

type OutputRecords struct {
	Records []OutputRecord
}

func (outputRecords *OutputRecords) ParseOutputsFromJson(jsonBody string) {
	modulesLocator := ModulesLocator
	jsonModules := gjson.Get(jsonBody, modulesLocator)

	if false == jsonModules.Exists() {
		return
	}

	outputRecords.Records = make([]OutputRecord, 0)

	jsonModules.ForEach(func(key, value gjson.Result) bool {
		records := make([]OutputRecord, 0)
		jsonOutputs := value.Get("outputs")

		jsonOutputs.ForEach(func(key, value gjson.Result) bool {
			sensitive := value.Get("sensitive")
			outputValueType := value.Get("type")
			outputValue := value.Get("value")

			outputRecord := OutputRecord{
				Name:      key.String(),
				Sensitive: sensitive.Bool(),
				Type:      outputValueType.String(),
				Value:     outputValue.String(),
			}

			records = append(records, outputRecord)

			return true
		})

		outputRecords.Records = append(outputRecords.Records, records...)

		return true
	})
}
