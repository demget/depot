// Package fs contains of Depot filesystem interfaces, where
// each of the subpackages is an implementation of these interfaces.
// It is being used by the Depot's client and server modules.
// The FS interface provides a basic set of functions to interact with
// any kind filesystem-like environment.
package fs

import (
	"io"
	"io/fs"
)

var (
	// ErrPermission is a shortcut to the std fs.ErrPermission.
	ErrPermission = fs.ErrPermission
)

// WalkFunc is a simplied walk function used for depot FS.
type WalkFunc func(path string, err error) error

// FS is a read-only fs used primarily by the server.
type FS interface {
	fs.FS

	// Meta returns the fs metadata.
	Meta() (Meta, error)
}

// WriteFS is a read-write fs used primarily by the client.
type WriteFS interface {
	FS

	// WriteFile writes the file using io.WriterTo.
	WriteFile(name string, w io.WriterTo) error
}

// Meta represents useful data of the fs.
type Meta struct {
	// Path is a list of all the files on the fs.
	Path []string
}
