package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

type Plugin struct {
	Git map[string]string `yaml:"git"`
}

var (
	PluginLocation      = new(Plugin)
	home                string
	packageFileLocation string
)

func init() {
	if err := setupPackageFileLocation(); err != nil {
		fmt.Println("can't get plugin package file.")
		os.Exit(1)
	}
}

func setupPluginLocaton() error {
	if _, err := os.Stat(packageFileLocation); os.IsNotExist(err) {
		fmt.Println("can't find local package file. fetch it from internet...")
		if err := FetchPluginPackage(); err != nil {
			fmt.Println("can't fetch plugin package file.")
			return err
		}
	}
	bs, err := ioutil.ReadFile(packageFileLocation)
	if err != nil {
		fmt.Println("can't read download package file.")
		return err
	}

	err = toml.Unmarshal(bs, PluginLocation)
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	return nil
}

func FetchPluginPackage() error {
	if err := createPackageFileDir(); err != nil {
		return err
	}

	resp, err := http.Get("https://xylonx.github.io/zshx/package.toml")
	if err != nil {
		fmt.Println("get package file error.")
		return err
	}
	defer resp.Body.Close()

	var dst *os.File
	if _, err := os.Stat(packageFileLocation); os.IsNotExist(err) {
		dst, err = os.Create(packageFileLocation)
		if err != nil {
			fmt.Println("create package.toml error")
			return err
		}
	} else {
		dst, err = os.Open(packageFileLocation)
		if err != nil {
			fmt.Println("open local package file error")
			return err
		}
	}
	defer dst.Close()

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		fmt.Println("download package to local failed")
		return err
	}

	return nil
}

func createPackageFileDir() error {
	// create diretory if it don't exist
	dir := home + "/.config/zshx"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// create file if does not exist
	filename := dir + "/package.toml"

	packageFileLocation = filename

	return nil
}

func setupPackageFileLocation() error {
	_home := os.Getenv("HOME")
	if _home == "" {
		return errors.New("can't get user's home dir from env $HOME")
	}
	home = _home
	packageFileLocation = home + "/.config/zshx/package.toml"
	return nil
}
