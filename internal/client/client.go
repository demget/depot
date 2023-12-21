package client

import (
	"os"
	"path/filepath"

	"github.com/pin/tftp"
)

type Client struct {
	dataClient *tftp.Client
	path       string
}

func New(addr, path string) (Client, error) {
	tftpc, err := tftp.NewClient(addr)
	if err != nil {
		return Client{}, err
	}

	return Client{tftpc, path}, nil
}

func (c Client) Recieve(remotePath string) error {
	wt, err := c.dataClient.Receive(remotePath, "octet")
	if err != nil {
		return err
	}

	path := c.path + remotePath
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = wt.WriteTo(file)
	if err != nil {
		return err
	}

	return nil
}
