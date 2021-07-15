package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xylonx/zshx/util"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch the upstream package",
	Long:  "force fetch the upstream package file to local",
	Run: func(cmd *cobra.Command, args []string) {
		if err := fetch(); err != nil {
			fmt.Println("fetch plugin package error: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}

func fetch() error {
	if err := util.FetchPluginPackage(); err != nil {
		return err
	}
	return nil
}
