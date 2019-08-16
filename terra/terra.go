package terra

import "io/ioutil"

type Recipe struct {
	Location string
	Body     string
}

func (recipe *Recipe) Create() error {
	fileBodyBytes, err := ioutil.ReadFile(recipe.Location)

	if nil != err {
		return err
	}

	recipe.Body = string(fileBodyBytes)

	return nil
}
