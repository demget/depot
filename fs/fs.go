package fs

import (
	"io"
	"io/fs"
)

var (
	ErrPermission = fs.ErrPermission
)

// Depot filesystem interfaces
// From caller perspective all the storage types
// like gdrive/telegram/osfs should behave just like
// normal filesystem, caller shouldn't be concerned
// with the implementation details of each filesystem

type WalkFunc func(path string, err error) error

// Read-only filesystem
type FS interface {
	Walk(fn WalkFunc) error
	Open(name string) (File, error)
}

// Read-Write filesystem
type RWFS interface {
	FS
	ReadFile(name string) ([]byte, error)
}

type ReadFileFS interface {
	FS
	WriteFile(name string, w io.WriterTo) error
}

type WriteFileFS interface {
	ReadFileFS
	RWFS
	WriteFile(name string, w io.WriterTo) error
}

// Read-Write file interface. It should behave exact
// like fs.File, but with Write() method
type File interface {
	fs.File
	Write([]byte) (n int, err error)
}
