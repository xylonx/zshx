package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xylonx/zshx/util"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all packages",
	PreRun: func(c *cobra.Command, args []string) {
		if err := util.Setup(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(c *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	fmt.Println("fetching plugin list...")

	for i := range util.PluginLocation.Git {
		fmt.Printf("%s[Git]:\n\t%s\n", i, util.PluginLocation.Git[i])
	}

	fmt.Println("fetching plugin done")
}
