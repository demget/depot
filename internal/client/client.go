package client

import (
	"bytes"
	"encoding/json"

	"github.com/demget/depot/fs"
	"github.com/demget/depot/internal"
	"github.com/demget/depot/pkg/netaddr"

	"github.com/pin/tftp"
)

type Client struct {
	fs   fs.WriteFS   // write-only filesystem
	tftp *tftp.Client // file transfer channel
}

func New(fs fs.WriteFS, addr string) (*Client, error) {
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
		tftp: client,
	}, nil
}

func (c *Client) Meta() (*fs.Meta, error) {
	wt, err := c.tftp.Receive("..", "octet")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err = wt.WriteTo(&buf); err != nil {
		return nil, err
	}

	var meta fs.Meta
	return &meta, json.NewDecoder(&buf).Decode(&meta)
}

func (c *Client) Read(name string) error {
	wt, err := c.tftp.Receive(name, "octet")
	if err != nil {
		return err
	}
	return c.fs.WriteFile(name, wt)
}
