package cli

import (
	"fmt"
	"os"

	"github.com/demget/depot/internal/server"
	"github.com/demget/depot/pkg/netaddr"

	"github.com/spf13/cobra"
)

func runServer(path, addr string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	host, port, err := netaddr.SplitHostPort(addr, server.DefaultPort)
	if err != nil {
		return err
	}

	s := server.New(host+":"+port, path)
	err = s.Start()
	if err != nil {
		return err
	}

	fmt.Printf("Depot is sharing '%s' on address %s.\n", path, addr)
	fmt.Printf("To connect, enter the command `depot client '%s'`.\n", addr)
	return nil
}

func NewCmdServer() *cobra.Command {
	var addr string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start depot fileserver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(args[0], addr)
		},
	}

	serverCmd.PersistentFlags().StringVar(&addr, "addr", "::1338", "Address for server to run")

	return serverCmd
}
