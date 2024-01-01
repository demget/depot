package tftpfs

import (
	"bytes"
	"io"
	"io/fs"
	"time"

	"github.com/pin/tftp"
)

type tftpFS struct {
	client *tftp.Client
}

type tftpFile struct {
	fileFS *tftpFS
	data   *bytes.Buffer
	size   int
	seek   int
	path   string
}

func (f *tftpFS) Open(path string) (fs.File, error) {
	if f == nil {
		return nil, fs.ErrInvalid
	}

	wt, err := f.client.Receive(path, "octet")
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	n, err := wt.WriteTo(buf)
	if err != nil {
		return nil, err
	}

	return &tftpFile{
		fileFS: f,
		data:   buf,
		size:   int(n),
		path:   path,
	}, nil
}

func (f *tftpFile) Stat() (fs.FileInfo, error) {
	if f == nil {
		return nil, fs.ErrInvalid
	}
	return f, nil
}

func (f *tftpFile) Read(dst []byte) (int, error) {
	if f.seek >= f.size {
		return 0, io.EOF
	}

	n := copy(dst, f.data.Bytes()[f.seek:f.size])
	f.seek += n
	return n, nil
}

func (f tftpFile) Close() error {
	return nil
}

func (f tftpFile) Name() string {
	return f.path
}

func (f tftpFile) Size() int64 {
	println(f.size)
	return int64(f.size)
}
func (f tftpFile) Mode() fs.FileMode {
	return 0400
}

func (f tftpFile) ModTime() time.Time {
	return time.Now()
}

func (f tftpFile) IsDir() bool {
	return false
}

func (f *tftpFile) Sys() any {
	return "tfptfs"
}
