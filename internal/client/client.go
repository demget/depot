package client

import (
	"os"
	"path/filepath"

	"github.com/pin/tftp"
)

type Client struct {
	client fs.Client
	output string
}

func New(c fs.Client, output string) (*Client, error) {
	return &Client{client: c, output: output}, nil
}

func (c *Client) Receive(remotePath string) error {
	file, err := c.client.Receive(remotePath)
	if err != nil {
		return err
	}

	fmt.Println("depot/client: received file:", file.FileName())
	return nil
}
