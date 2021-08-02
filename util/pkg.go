package util

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
)

type Plugin struct {
	Git map[string]string `yaml:"git"`
}

var PluginLocation = new(Plugin)

func setupPluginLocaton() error {
	resp, err := http.Get("https://raw.githubusercontent.com/xylon/zshx/master/package.toml")
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	err = toml.Unmarshal(bs, PluginLocation)
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	return nil
}
