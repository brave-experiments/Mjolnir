package terra

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

const (
	ModulesLocator = "modules"
)

var (
	SupportedType = "string"
)

type OutputRecord struct {
	Name      string
	Sensitive bool        `json:"sensitive"`
	Type      string      `json:"type"`
	Value     interface{} `json:"value"`
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
		err := json.Unmarshal([]byte(jsonOutputs.Raw), &records)

		// This is key => value not []value!

		if nil != err {
			fmt.Println("Error!: ", err)
			return false
		}

		outputRecords.Records = append(outputRecords.Records, records...)

		return true
	})

	fmt.Println("records!", outputRecords.Records)
}
