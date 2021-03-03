package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compomize",
	Short: "compomize is the utility which Convert k8s Manifest to docker-compose.yaml",
}

func printError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
}

func exitError(msg interface{}) {
	printError(msg)
	os.Exit(1)
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		rootCmd.Help()
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
