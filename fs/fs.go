package fs

import (
	"io"
	"io/fs"
)

var (
	ErrPermission = fs.ErrPermission
)

type FS interface {
	fs.FS

	WriteFile(name string, w io.WriterTo) error
}
