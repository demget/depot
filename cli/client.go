package cli

import (
	"fmt"
	"os"

	"github.com/demget/depot/internal/server"
	"github.com/demget/depot/pkg/netaddr"

	"github.com/spf13/cobra"
)

func runClient(path, addr string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	_, _, err = netaddr.SplitHostPort(addr, server.DefaultPort)
	if err != nil {
		return err
	}

	fmt.Printf("Depot client connected to %s successfully!\n", addr)
	fmt.Printf("Synchronizing files in %s directory\n", path)
	return nil
}

func NewCmdClient() *cobra.Command {
	var path string

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "Connect to depot fileserver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClient(path, args[0])
		},
	}

	clientCmd.PersistentFlags().StringVarP(&path, "output", "o", ".", "Sync directory path")

	return clientCmd
}
