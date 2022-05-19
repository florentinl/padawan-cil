package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "padawan",
	Short: "Padawan is an awesome cli to interact with the padawan container service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Padawan is an awesome cli to interract with the padawan container service")
	},
}

func init() {
	imgCmd.AddCommand(imgAddCmd)
	imgCmd.AddCommand(imgDelCmd)
	imgCmd.AddCommand(imgLsCmd)
	imgCmd.AddCommand(imgGetCmd)
	imgCmd.AddCommand(imgSetCmd)

	ctrCmd.AddCommand(ctrLsCmd)
	ctrCmd.AddCommand(ctrRunCmd)
	ctrCmd.AddCommand(ctrDelCmd)
	ctrCmd.AddCommand(ctrGetCmd)

	rootCmd.AddCommand(ctrCmd)
	rootCmd.AddCommand(imgCmd)
	rootCmd.AddCommand(loginCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
