package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1"

func Run() {
	var (
		rootCmd = &cobra.Command{
			Use:   "depot",
			Short: "📦 Depot v" + Version,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
			},
		}

		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version number of depot",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Depot", "v"+Version)
			},
		}
	)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(NewCmdServer())
	rootCmd.AddCommand(NewCmdClient())

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
