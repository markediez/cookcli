package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ReadCommand struct {}

func (*ReadCommand) New() *cobra.Command {
	return &cobra.Command {
		Use: "read",
		Short: "read something",
		Long: "read something long desc",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("TODO: read something")

			return nil
		},
	}
}