package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Plugin struct {
	Git map[string]string `yaml:"git"`
}

var PluginLocation = new(Plugin)

func setupPluginLocaton() error {
	file, err := os.Open("./package.yaml")
	if err != nil {
		fmt.Println("get config file error")
		return err
	}
	bs, _ := ioutil.ReadAll(file)
	if err := yaml.Unmarshal(bs, PluginLocation); err != nil {
		fmt.Println("unmarshal error.")
		return err
	}
	return nil
}
