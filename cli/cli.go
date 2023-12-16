package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type cobraCmdHandler = func(cmd *cobra.Command, args []string) error

func mockStartServer(addr, path string) error {
	fmt.Printf("Depot is sharing '%s' on address %s.\n", path, addr)
	fmt.Printf("To connect, enter the command `depot client --addr %s`.\n", addr)

	return nil
}

func mockStartClient(addr, path string) error {
	fmt.Printf("Depot client connected to %s successfully!\n", addr)

	return nil
}

func verboseCliInfo(target cobraCmdHandler) cobraCmdHandler {
	return func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Command: %s\nArgs: %v\n\n", cmd.Name(), args)
		return target(cmd, args)
	}
}

func Run() {
	var addr string

	var (
		rootCmd = &cobra.Command{
			Use:   "depot",
			Short: "depot",
			RunE: verboseCliInfo(func(cmd *cobra.Command, args []string) error {
				println("?")
				return nil
			}),
		}

		versionCmd = &cobra.Command{
			Use:   "version",
			Short: "Print the version number of depot",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("depot v0.0")
			},
		}

		serverCmd = &cobra.Command{
			Use:   "server",
			Short: "Start depot fileserver",
			Args:  cobra.ExactArgs(1),
			RunE: verboseCliInfo(func(cmd *cobra.Command, args []string) error {
				return mockStartServer(addr, args[0])
			}),
		}

		clientCmd = &cobra.Command{
			Use:   "client",
			Short: "Connect to depot fileserver",
			Args:  cobra.ExactArgs(1),
			RunE: verboseCliInfo(func(cmd *cobra.Command, args []string) error {
				return mockStartClient(addr, args[0])
			}),
		}
	)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(clientCmd)

	serverCmd.PersistentFlags().StringVar(&addr, "addr", "::1338", "Address for server to run")
	clientCmd.PersistentFlags().StringVar(&addr, "addr", "::1338", "Address of server")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
