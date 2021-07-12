// TODO: install new plugin
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/xylonx/zshx/util"
)

type replaceFunc func(str string) bool

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
		err := install()
		if err != nil {
			os.Exit(1)
		} else {
			fmt.Println("install done.")
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

// TODO: develop the service logic
func install() error {
	regex, err := regexp.Compile(`^\splugins\s=\s\((.*)\)$`)
	if err != nil {
		fmt.Println("compile regex failed. err: ", err)
		return err
	}

	file, err := os.Open(zshrcPath)
	if err != nil {
		fmt.Println("open zshrc file failed. err: ", err)
		return err
	}
	defer file.Close()

	tmpfile, err := os.Create(zshrcPath + ".zsh.tmp")
	if err != nil {
		fmt.Println("create zshrc tmp file failed. err: ", err)
		return err
	}
	// TODO: remove tmp file
	defer func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	err = replaceByLine(file, tmpfile, func(str string) bool {
		matches := regex.FindStringSubmatch(str)
		fmt.Println(matches)
		return true
	})
	if err != nil {
		fmt.Println("replace file by line failed. err: ", err)
		return err
	}

	return nil
}

func replaceByLine(r io.Reader, w io.Writer, fn replaceFunc) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		fn(line)
		if _, err := w.Write([]byte(line + "\n")); err != nil {
			fmt.Println("write to tmp file failed. err: ", err)
			return err
		}
	}
	return sc.Err()
}
