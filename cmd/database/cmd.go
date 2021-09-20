package database

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "db",
		Short: "db command",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCmd(args)
			fmt.Println(err)
		},
	}
}
