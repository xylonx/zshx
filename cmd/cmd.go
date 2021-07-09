package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xylonx/zshx/util"
)

const (
	DEFAULT_ZSHRC_PATH = "$HOME/.zshrc"
)

var (
	rootCmd = &cobra.Command{
		Use:   "zsh-plugin-helper",
		Short: "a command tool to install oh my zsh plugins",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := util.Setup(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("install missing zsh plugin...")
			if err := update(); err != nil {
				log.Fatalln(err)
			}
		},
	}

	zshrcPath string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&zshrcPath, "zshrc", "z", DEFAULT_ZSHRC_PATH, "specify the path of the .zshrc")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("error occur: ", err)
		os.Exit(1)
	}
}

func update() error {
	if zshrcPath == DEFAULT_ZSHRC_PATH {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		zshrcPath = home + "/.zshrc"
	}

	s, err := util.GetPluginFromZshrc(zshrcPath)
	if err != nil {
		return err
	}

	fmt.Printf("detect your plugins: %s\n", strings.Join(s, " "))

	toInstall, err := util.DetectNotInstalledPlugin(s)
	if err != nil {
		return err
	}

	if len(toInstall) == 0 {
		fmt.Println("all plugin have been installed")
		return nil
	}

	// TODO: install missing plugins
	for i := range toInstall {
		err := util.InstallPluginByGit(toInstall[i], util.PluginLocation.Git[toInstall[i]])
		if err != nil {
			fmt.Printf("install missing plugins %s error: %v\n", toInstall[i], err)
		} else {
			fmt.Printf("install missing plugins %s successfully.\n", toInstall[i])
		}
	}

	return nil
}
