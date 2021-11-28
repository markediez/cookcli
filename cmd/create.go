package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CreateCommand struct {}

func (*CreateCommand) New() *cobra.Command {
	return &cobra.Command {
		Use: "create",
		Short: "create something",
		Long: "create something long desc",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO: Create something")

			return nil
		},
	}
}