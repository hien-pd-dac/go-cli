package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli [OPTIONS] [COMMANDS]",
	Short: "go-cli is a sample command.",
	Long:  `go-cli is a sample command.`,
	// Args:  cobra.MinimumNArgs(1),
	// RunE runs with error handling.
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("you must use at least one argument.\n")
		}
		fmt.Printf("Hello world! you have typed with more than zero argument: `go-cli %s`.\n", strings.Join(args, " "))
		return nil
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// fmt.Printf("hello world!")
	// },
}

// Execute runs the commander.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// init initilizes the application configuration and defines flags.
// init will been called automatically by cobra before Execute() function be processed.
func init() {
	initVersionCmd()
	initGetUsersCmd()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getUsersCmd)
}
