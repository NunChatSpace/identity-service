package server

import (
	"github.com/NunChatSpace/identity-service/http"
	"github.com/spf13/cobra"
)

func GetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Listen and serve",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			http.GetServer().Run()
		},
	}
}
