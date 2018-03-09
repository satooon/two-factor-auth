package yaml

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func Load(filename string, out interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}
