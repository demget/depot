package cli

import (
	"fmt"
	"os"

	"github.com/demget/depot/fs/osfs"
	"github.com/demget/depot/internal/client"

	"github.com/spf13/cobra"
)

func runClient(addr, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	c, err := client.New(osfs.New(path), addr)
	if err != nil {
		return err
	}
	if err := c.Read(path); err != nil {
		return err
	}

	fmt.Printf("Depot client connected to %s successfully!\n", addr)
	fmt.Printf("Synchronizing files in %s directory.\n", path)

	return nil
}

func NewCmdClient() *cobra.Command {
	var path string

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "Connect to depot fileserver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClient(args[0], path)
		},
	}

	clientCmd.PersistentFlags().StringVarP(&path, "output", "o", ".", "Sync directory path")

	return clientCmd
}
