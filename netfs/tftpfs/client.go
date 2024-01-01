package tftpfs

import (
	"io/fs"

	"github.com/demget/depot/pkg/netaddr"
	"github.com/pin/tftp"
)

type Client struct {
	addr  string
	files tftpFS
}

func New(addr string) *Client {
	return &Client{addr: addr}
}

func (c *Client) Connect() (fs.FS, error) {
	var err error

	host, port, err := netaddr.SplitHostPort(c.addr, defaultPort)
	if err != nil {
		return nil, err
	}

	c.files.client, err = tftp.NewClient(host + ":" + port)
	if err != nil {
		return nil, err
	}

	return &c.files, nil
}

func (c *Client) Disconnect() error {
	return nil
}
