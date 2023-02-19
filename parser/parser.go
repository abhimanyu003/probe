package parser

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func ParseYaml(data []byte) ([]byte, error) {
	m := make(map[string]interface{})
	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	j, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return j, nil
}
