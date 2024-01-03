package cli

import (
	"fmt"
	"os"

	"github.com/demget/depot/fs/osfs"
	"github.com/demget/depot/internal/server"

	"github.com/spf13/cobra"
)

func runServer(addr, path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	s, err := server.New(osfs.New(path), addr)
	if err != nil {
		return err
	}

	fmt.Printf("Depot is sharing '%s' on address %s.\n", path, addr)
	fmt.Printf("To connect, enter the command `depot client %s`.\n", addr)

	return s.Start()
}

func NewCmdServer() *cobra.Command {
	var addr string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start depot fileserver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(addr, args[0])
		},
	}

	serverCmd.PersistentFlags().StringVar(&addr, "addr", "::1338", "Address for server to run")

	return serverCmd
}
