package tftp

import (
	"os"
	"path/filepath"

	"github.com/pin/tftp"
)

type Client struct {
	*tftp.Client
	fs fs.FS
}

func NewClient(addr string, fs fs.FS) (*Client, error) {
	tftpc, err := tftp.NewClient(addr)
	if err != nil {
		return nil, err
	}

	return &Client{Client: tftpc, fs: fs}, nil
}

func (c *Client) Receive(path string) (fs.File, error) {
	wt, err := c.Receive(path, "octet")
	if err != nil {
		return err
	}

	return c.fs.Write(path, wt)
}
