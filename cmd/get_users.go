package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UserID int64
var Email string

var getUsersCmd = &cobra.Command{
	Use:   "get-users",
	Short: "go-cli get-users gets user's infos.",
	Long:  `go-cli get-users gets user's infos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Get users's data by: --user_id=%d --email=%s\n", UserID, Email)
	},
}

func initGetUsersCmd() {
	//note: if you set shorthand for `user_id` as `id`, this will cause the error: `panic: "id" shorthand is more than one ASCII character`
	getUsersCmd.Flags().Int64VarP(&UserID, "user_id", "i", 0, "Filters users's data by user_id.")
	getUsersCmd.Flags().StringVarP(&Email, "email", "e", "", "Filters users'data by email.")
	return
}
