package internal

import (
	"errors"
	"io/fs"
	"net/url"
	"os"
	"strings"

	"github.com/demget/depot/netfs"
	"github.com/demget/depot/netfs/tftpfs"
)

type Client struct {
	client netfs.Client
	src    fs.FS
}

var (
	ErrUnknownScheme = errors.New("unknown scheme")
	ErrNotConnected  = errors.New("client isn't connected")
)

func NewClient(addr string) (*Client, error) {
	c, err := chooseBackend(addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: c,
	}, nil
}

func (c *Client) Connect() error {
	src, err := c.client.Connect()
	if err != nil {
		return err
	}

	c.src = src
	return nil
}

func (c *Client) Sync(dst string) error {
	if c.src == nil {
		return ErrNotConnected
	}

	data, err := fs.ReadFile(c.src, "test.txt")
	if err != nil {
		return err
	}

	file, err := os.OpenFile(dst+"/test.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func chooseBackend(uri string) (netfs.Client, error) {
	u, err := url.Parse(uri)
	fsBackend := netfs.DefaultFS
	if err != nil {
		if !strings.Contains(err.Error(), "missing protocol scheme") {
			return nil, err
		}
	} else {
		fsBackend = u.Scheme
	}

	switch fsBackend {
	case "tftp":
		return tftpfs.New(u.Hostname()), nil
	default:
		return nil, ErrUnknownScheme
	}
}
