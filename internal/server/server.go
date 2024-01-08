package server

import (
	"fmt"
	"io"
	"strings"

	"github.com/demget/depot/fs"

	"github.com/pin/tftp"
)

type Server struct {
	fs fs.FS // read-only filesystem
	//	http *http.Server // communication channel
	tftp *tftp.Server // file transfer channel
	addr string
}

func New(fs fs.FS, addr string) (*Server, error) {
	s := &Server{fs: fs, addr: addr}
	s.tftp = tftp.NewServer(s.readHandler, s.writeHandler)
	return s, nil
}

func (s *Server) Start() error {
	return s.tftp.ListenAndServe(s.addr)
}

func (s *Server) Stop() error {
	s.tftp.Shutdown()
	return nil
}

func (s *Server) readHandler(name string, rf io.ReaderFrom) error {
	// handle meta files in ".." directory
	// since FS is by default read-only,
	// we can't write real files here
	// it's also safe because FS wouldn't
	// permit to read anything from above root dir.
	if name == ".." {
		var files string

		s.fs.Walk(func(path string, err error) error {
			files = fmt.Sprintln(files, path)
			return nil
		})

		_, err := rf.ReadFrom(strings.NewReader(files))
		return err
	}

	file, err := s.fs.Open(name)
	if err != nil {
		return err
	}
	_, err = rf.ReadFrom(file)
	return err
}

func (s *Server) writeHandler(name string, wt io.WriterTo) error {
	return fs.ErrPermission
}
