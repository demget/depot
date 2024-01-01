package cli

import (
	"fmt"
	"os"

	"github.com/demget/depot/internal"
	"github.com/spf13/cobra"
)

func runClient(path, addr string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	c, err := internal.NewClient(addr)
	if err != nil {
		return err
	}

	if err = c.Connect(); err != nil {
		return err
	}

	if err = c.Sync(path); err != nil {
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
