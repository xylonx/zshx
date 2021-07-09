// TODO: install new plugin
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xylonx/zshx/util"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install missing plugins",
	Long:  "install missing plugins from zshrc plugin",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := util.Setup(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start installing...")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
