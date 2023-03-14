package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 收益曲线、
//最大回撤、年化收益率、夏普比

// 根命令
var rootCmd = &cobra.Command{
	Use:   "fbinance",
	Short: "一个命令行应用程序",
	Long:  "这是一个收益计算工具",
}

// 子命令
var (
	serverCmd = &cobra.Command{
		Use:   "makeline",
		Short: "生成收益率曲线",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: 在此处添加启动服务器的代码
		},
	}

	dbCmd = &cobra.Command{
		Use:   "maxdrawdown",
		Short: "最大回撤",
		Long:  "收益最大回撤",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: 在此处添加启动服务器的代码
		},
	}

	// dbCmd 的子命令
	dbMigrateCmd = &cobra.Command{
		Use:   "年化收益率",
		Short: "执行数据库迁移",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: 在此处添加数据库迁移的代码
		},
	}
)

func init() {
	// 添加子命令
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(dbCmd)

	// 添加 dbCmd 的子命令
	dbCmd.AddCommand(dbMigrateCmd)
}

func main() {
	// 解析命令行参数并执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
