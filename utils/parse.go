package utils

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strings"
)

func ParseSelector(path string) (config map[string]interface{}) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		return ParseYaml(file)
	} else if strings.HasSuffix(path, ".json") {
		return ParseJson(file)
	} else if strings.HasSuffix(path, ".toml") {
		return ParseToml(file)
	}
	return map[string]interface{}{}
}

func ParseYaml(b []byte) (config map[string]interface{}) {
	config = make(map[string]interface{})
	err := yaml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return
}

func ParseJson(b []byte) (config map[string]interface{}) {
	config = make(map[string]interface{})
	err := json.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return
}

func ParseToml(b []byte) (config map[string]interface{}) {
	config = make(map[string]interface{})
	err := toml.Unmarshal(b, &config)
	if err != nil {
		panic(err)
	}
	return
}
