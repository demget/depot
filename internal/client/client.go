package client

import (
	"net/http"

	"github.com/demget/depot/fs"
	"github.com/demget/depot/internal"
	"github.com/demget/depot/pkg/netaddr"

	"github.com/pin/tftp"
)

type Client struct {
	fs   fs.FS        // write-only filesystem
	http *http.Client // communication channel
	tftp *tftp.Client // file transfer channel
}

func New(fs fs.FS, addr string) (*Client, error) {
	host, port, err := netaddr.SplitHostPort(addr, internal.DefaultPort)
	if err != nil {
		return nil, err
	}

	client, err := tftp.NewClient(host + ":" + port)
	if err != nil {
		return nil, err
	}

	return &Client{
		fs:   fs,
		http: &http.Client{},
		tftp: client,
	}, nil
}

func (c *Client) Read(name string) error {
	wt, err := c.tftp.Receive(name, "octet")
	if err != nil {
		return err
	}
	return c.fs.WriteFile(name, wt)
}
