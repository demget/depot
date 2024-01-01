package netfs

import "io/fs"

// Server represents a server that shares the FS with its clients.
// Server should implement at least read only file transfering.
type Server interface {
	Start(root fs.FS) error
	Stop() error
}

// Client represent client-server connection as fs.FS interface.
type Client interface {
	Connect() (fs.FS, error)
	Disconnect() error
}
