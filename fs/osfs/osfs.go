package osfs

import (
	"io"
	"io/fs"
	"os"
)

func New(dir string) *FS {
	return &FS{FS: os.DirFS(dir)}
}

type FS struct {
	fs.FS
}

func (fs FS) WriteFile(name string, w io.WriterTo) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = w.WriteTo(file)
	return err
}
