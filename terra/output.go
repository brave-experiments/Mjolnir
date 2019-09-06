package terra

import (
	"bytes"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

const (
	ModulesLocator        = "modules"
	ColorizedOutputPrefix = "[reset][bold][green]\nOutputs:\n\n"
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

func (outputRecords *OutputRecords) FromJsonAsString(jsonBody string, includeHeader bool) string {
	outputRecords.parseOutputsFromJson(jsonBody)
	outputs := outputRecords.Records

	if len(outputs) < 1 {
		return ""
	}

	outputBuf := new(bytes.Buffer)

	if includeHeader {
		outputBuf.WriteString(ColorizedOutputPrefix)
	}

	for key := range outputs {
		outputRecord := outputs[key]
		outputRecordName := outputRecord.Name

		if outputRecord.Sensitive {
			outputBuf.WriteString(fmt.Sprintf("%s = <sensitive>\n", outputRecordName))
			continue
		}

		result := outputRecord.Value

		outputBuf.WriteString(fmt.Sprintf("%s = %s\n", outputRecordName, result))
	}

	return strings.TrimSpace(outputBuf.String())
}

func (outputRecords *OutputRecords) parseOutputsFromJson(jsonBody string) {
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
