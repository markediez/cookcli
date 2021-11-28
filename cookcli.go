package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/markediez/cookcli/cmd"
)

func main() {
	rootCmd := &cobra.Command {
		Use: "cookcli",
		Short: "CLI interface for cook api",
	}

	rootCmd.AddCommand((&cmd.ListCommand{}).New())
	rootCmd.AddCommand((&cmd.ReadCommand{}).New())
	rootCmd.AddCommand((&cmd.CreateCommand{}).New())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
