package osfs

import (
	"io"
	"os"
	iofs "io/fs"

	"github.com/demget/depot/fs"
)

type FS struct {
	iofs.FS
}

func New(dir string) *FS {
	return &FS{FS: os.DirFS(dir)}
}

func (f *FS) WriteFile(name string, wt io.WriterTo) error {
	return os.WriteFile(name, wt, 0644)
}

func (f *FS) Meta() (fs.Meta, error) {
	return fs.Meta{
		Path: f.metaPath(),
	}, nil
}

func (f *FS) metaPath() (path []string) {
	walk := func(p string, d iofs.DirEntry, _ error) error {
		if !d.IsDir() {
			path = append(path, p)
		}
		return nil
	}
	iofs.WalkDir(f, ".", walk)
	return path
}
