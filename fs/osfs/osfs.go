// Package osfs provides a way to serve your OS filesystem.
// It implements the depot/fs.
package osfs

import (
	"bytes"
	"io"
	iofs "io/fs"
	"os"
	"path"

	"github.com/demget/depot/fs"
)

type FS struct {
	iofs.FS
}

func New(dir string) *FS {
	return &FS{FS: os.DirFS(dir)}
}

func (f *FS) WriteFile(name string, wt io.WriterTo) error {
	var buf bytes.Buffer
	if _, err := wt.WriteTo(&buf); err != nil {
		return err
	}

	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.MkdirAll(path.Dir(name), 0755); err != nil {
			return err
		}
	}

	return os.WriteFile(name, buf.Bytes(), 0644)
}

func (f *FS) Meta() (fs.Meta, error) {
	files, err := f.metaPath()
	if err != nil {
		return fs.Meta{}, err
	}
	return fs.Meta{
		Files: files,
	}, nil
}

func (f *FS) metaPath() (files []string, err error) {
	walk := func(p string, d iofs.DirEntry, _ error) error {
		if !d.IsDir() {
			files = append(files, p)
		}
		return nil
	}
	err = iofs.WalkDir(f, ".", walk)
	return files, err
}
