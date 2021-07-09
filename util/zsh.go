package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	PLGIN_PATTERN = "^plugins\\s*=\\s*\\((.*)\\)$"
)

var (
	PLUGIN_HOLDER_DEFAULT string
	PLUGIN_HOLDER_CUSTOM  string
)

func setupZsh() error {
	if os.Getenv("ZSH") == "" {
		return errors.New("can't get ZSH environment variable")
	}
	PLUGIN_HOLDER_DEFAULT = os.Getenv("ZSH") + "/plugins/"
	PLUGIN_HOLDER_CUSTOM = os.Getenv("ZSH") + "/custom/plugins/"
	return nil
}

func GetPluginFromZshrc(zshrcPath string) ([]string, error) {
	file, err := os.Open(zshrcPath)
	if err != nil {
		return nil, err
	}

	regex, err := regexp.Compile(PLGIN_PATTERN)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		match := regex.FindStringSubmatch(txt)
		if len(match) == 2 {
			plugins := strings.Fields(strings.TrimSpace(match[1]))
			return plugins, nil
		}
	}

	return nil, errors.New("not detect the plugin")
}

func DetectNotInstalledPlugin(plugin []string) ([]string, error) {
	defaultEntries, err := os.ReadDir(PLUGIN_HOLDER_DEFAULT)
	if err != nil {
		return nil, err
	}
	customEntries, err := os.ReadDir(PLUGIN_HOLDER_CUSTOM)
	if err != nil {
		return nil, err
	}
	missing := FindMissing(plugin, defaultEntries)
	missing = FindMissing(missing, customEntries)
	return missing, nil
}

// Default install location: custom/plugin
func InstallPluginByGit(pluginname, gitPath string) error {
	var out bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &out)
	cmd := exec.Command("git", "clone", gitPath, PLUGIN_HOLDER_CUSTOM+pluginname)
	cmd.Stdout = mw
	cmd.Stderr = mw
	// go func() { fmt.Println(mw) }()
	return cmd.Run()
}
