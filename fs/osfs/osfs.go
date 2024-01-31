// Package osfs implements Depot FS using os.DirFS
// for RO access plus os.WriteFile
package osfs

import (
	"bytes"
	"io"
	iofs "io/fs"
	"os"
	"path"
	"strings"

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
	err := os.WriteFile(name, buf.Bytes(), 0644)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		err = os.MkdirAll(path.Dir(name), 0755)
		if err != nil {
			return err
		}
		return os.WriteFile(name, buf.Bytes(), 0644)
	}

	return err
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
