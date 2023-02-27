/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "github.com/spf13/cobra"

var (
	varExeName       string
	varStringGo      string
	varStringBase    string
	varIntPort       int
	varStringHome    string
	varStringRemote  string
	varStringBranch  string
	varStringVersion string
	varStringTZ      string

	Cmd = &cobra.Command{
		Use:   "docker",
		Short: "生成Dockerfile文件",
		RunE:  dockerCommand,
	}
)
