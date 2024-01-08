package osfs

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	iofs "io/fs"

	"github.com/demget/depot/fs"
)

func New(dir string) fs.WriteFileFS {
	return &FS{root: dir}
}

type FS struct {
	root string
}

type File = os.File

func (fsys *FS) Open(name string) (fs.File, error) {
	path := fsys.root + "/" + name
	if err := fsys.checkRoot(path); err != nil {
		return nil, err
	}

	return os.Open(path)
}

func (fsys *FS) Walk(fn fs.WalkFunc) error {
	return filepath.WalkDir(fsys.root, func(path string, d iofs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		return fn(path, err)
	})
}

func (fsys *FS) OpenFile(name string) (*File, error) {
	path := fsys.root + "/" + name
	if err := fsys.checkRoot(path); err != nil {
		return nil, err
	}

	return os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
}

func (fsys *FS) ReadFile(name string) ([]byte, error) {
	path := fsys.root + "/" + name
	if err := fsys.checkRoot(path); err != nil {
		return nil, err
	}

	return os.ReadFile(path)
}

func (fsys *FS) WriteFile(name string, w io.WriterTo) error {
	if err := fsys.checkRoot(name); err != nil {
		return err
	}

	var buf bytes.Buffer
	w.WriteTo(&buf)
	return os.WriteFile(name, buf.Bytes(), 0644)
}

func (fsys *FS) checkRoot(path string) error {
	rootAbs, err := filepath.Abs(fsys.root)
	if err != nil {
		return err
	}

	pathAbs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(pathAbs, rootAbs) {
		return fs.ErrPermission
	}

	return nil
}
