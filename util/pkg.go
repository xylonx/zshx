package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Plugin struct {
	Git map[string]string `yaml:"git"`
}

var PluginLocation = new(Plugin)

func setupPluginLocaton() error {
	resp, err := http.Get("https://xylonx.github.io/zshx/package.yaml")
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	if err := yaml.Unmarshal(bs, PluginLocation); err != nil {
		fmt.Println("unmarshal error.")
		return err
	}
	return nil
}
