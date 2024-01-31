package server

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/demget/depot/fs"

	"github.com/pin/tftp"
)

type Server struct {
	fs   fs.FS        // read-only filesystem
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

func (s *Server) readHandler(name string, rf io.ReaderFrom) (err error) {
	var r io.Reader

	// Handle meta files in ".." directory. Since FS is read-only by default,
	// we can't write real files here, it's also safe because FS wouldn't
	// permit to read anything from the dir above root.
	if name == ".." {
		meta, err := s.fs.Meta()
		if err != nil {
			return err
		}

		data, err := json.Marshal(meta)
		if err != nil {
			return err
		}

		r = bytes.NewBuffer(data)
	} else {
		r, err = s.fs.Open(name)
		if err != nil {
			return err
		}
	}

	_, err = rf.ReadFrom(r)
	return err
}

func (s *Server) writeHandler(_ string, _ io.WriterTo) error {
	// Write is not allowed for the server.
	return fs.ErrPermission
}
