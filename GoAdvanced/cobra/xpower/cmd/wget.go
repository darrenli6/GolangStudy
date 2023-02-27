/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	output string
)

// wgetCmd represents the wget command
var wgetCmd = &cobra.Command{
	Use: "wget",

	/*
		cobra内置的参数验证也是比较多，
		NoArgs、OnlyValidArgs、MinimumNArgs、MaximumNArgs
		等等可翻阅源码args.go，可以满足基本使用，
		如果有自己的特殊要求可以通过解析arg来实现。

	*/
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wget called")
		Download(args[0], output)
	},
}

func Download(url string, path string) {

	out, err := os.Create(path)
	check(err)
	defer out.Close()

	res, err := http.Get(url)
	check(err)
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	check(err)
	fmt.Println("save as " + path)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(wgetCmd)

	/*
		flag包含局部和全局两种，全局flag在父命令定义后子命令也会生效，
		而局部flag则在哪定义就在哪生效。

		如上面的局部flag，我们在wgetCmd中定义的flag只有wget这个子命令能用

	*/

	wgetCmd.Flags().StringVarP(&output, "output", "o", "", "outfile")
	wgetCmd.MarkFlagRequired("output")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wgetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wgetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
