package terra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

const (
	ModulesLocator        = "modules"
	PrivateKeyLocator     = "tls_private_key.ssh"
	BastionKeyLocator     = "aws_iam_instance_profile.bastion"
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

type keyPair struct {
	DeployName        string
	PrivateKey        string `json:"private_key_pem"`
	PublicKey         string `json:"public_key_pem"`
	RsaBits           string `json:"rsa_bits"`
	Algorithm         string `json:"algorithm"`
	Id                string `json:"id"`
	EcdsaCurve        string `json:"ecdsa_curve"`
	PublicFingerprint string `json:"public_key_fingerprint_md5"`
	OpenSsh           string `json:"public_key_openssh"`
	privateKeyFile    File
	publicKeyFile     File
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

func (keyPair *keyPair) FromJson(jsonBody string) {
	modulesLocator := ModulesLocator
	jsonModules := gjson.Get(jsonBody, modulesLocator)

	if false == jsonModules.Exists() {
		return
	}

	jsonModules.ForEach(func(key, value gjson.Result) bool {
		jsonResources := value.Get("resources")
		shouldIterate := true

		if false == jsonResources.Exists() {
			return false
		}

		jsonResources.ForEach(func(key, value gjson.Result) bool {
			keyPair.mapName(key.String(), value)
			keyPair.marshalKeyPair(key.String(), value, &shouldIterate)

			return shouldIterate
		})

		return shouldIterate
	})
}

func (outputRecords *OutputRecords) parseOutputsFromJson(jsonBody string) {
	modulesLocator := ModulesLocator
	jsonModules := gjson.Get(jsonBody, modulesLocator)

	if false == jsonModules.Exists() {
		return
	}

	jsonModules.ForEach(func(key, value gjson.Result) bool {
		jsonOutputs := value.Get("outputs")
		jsonOutputs.ForEach(outputRecords.mapRecords)

		return true
	})
}

func (outputRecords *OutputRecords) mapRecords(key, value gjson.Result) (shouldIterate bool) {
	if nil == outputRecords.Records {
		outputRecords.Records = make([]OutputRecord, 0)
	}

	sensitive := value.Get("sensitive")
	outputValueType := value.Get("type")
	outputValue := value.Get("value")

	outputRecord := OutputRecord{
		Name:      key.String(),
		Sensitive: sensitive.Bool(),
		Type:      outputValueType.String(),
		Value:     outputValue.String(),
	}

	outputRecords.Records = append(outputRecords.Records, outputRecord)

	return true
}

func (keyPair *keyPair) marshalKeyPair(key string, value gjson.Result, shouldIterate *bool) {
	if PrivateKeyLocator != key {
		return
	}

	jsonKeyPair := value.Get("primary.attributes")

	if false == jsonKeyPair.Exists() {
		return
	}

	err := json.Unmarshal([]byte(jsonKeyPair.Raw), &keyPair)

	if nil != err {
		*shouldIterate = false
	}

	return
}

func (keyPair *keyPair) mapName(key string, value gjson.Result) {
	if BastionKeyLocator != key {
		return
	}

	deploymentName := value.Get("primary.id")

	if false == deploymentName.Exists() {
		return
	}

	keyPair.DeployName = deploymentName.String()
}
