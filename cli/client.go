package cli

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/demget/depot/netfs/tftpfs"
	"github.com/spf13/cobra"
)

func runClient(path, addr string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	s := tftpfs.NewClient(addr)
	fsys, err := s.Connect()
	if err != nil {
		return err
	}

	buf, err := fs.ReadFile(fsys, "cli.go")
	if err != nil {
		return err
	}

	err = os.WriteFile("test.txt", buf, 0755)
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
