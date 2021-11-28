package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ListCommand struct {}

func (*ListCommand) New() *cobra.Command {
	return &cobra.Command {
		Use: "list",
		Short: "print a list of stuff",
		Long: "print a list of paginated stuff",
		RunE: func(cmd *cobra.Command, arg []string) error {
			fmt.Println("TODO: Query API")
			fmt.Println("Sample Output:")
			fmt.Println("Butter Pie Crust")
			fmt.Println("Honey Drop Biscuits")
			fmt.Println("Pumpkin Pie")

			return nil
		},
	}
}