package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(sub1, sub2)
	sub1.AddCommand(sub2)
}

func main() {
	root.Execute()
}

var root = &cobra.Command{
	Use: "root",

	// Persistent Run
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun in root")
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var config string

func init() {
	// Persistent Flag
	root.PersistentFlags().StringVarP(&config, "config", "c", "~/.config.json", "配置文件")
}

var sub1 = &cobra.Command{
	Use: "sub1",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun in sub1")
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	sub1.PersistentFlags().StringVarP(&config, "config", "c", "$HOME/.config.json", "配置文件")
}

var sub2 = &cobra.Command{
	Use: "sub2",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
