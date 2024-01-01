package tftpfs

import (
	"io"
	"io/fs"

	"github.com/demget/depot/netfs"
	"github.com/pin/tftp"
)

type Server struct {
	server *tftp.Server
	addr   string
	root   fs.FS
}

const defaultPort = "1338"

func NewServer(addr string) netfs.Server {
	var s Server

	s.server = tftp.NewServer(s.readHandler, s.writeHandler)
	s.addr = addr
	return &s
}

func (s *Server) Start(root fs.FS) error {
	s.root = root
	return s.server.ListenAndServe(s.addr)
}

func (s *Server) Stop() error {
	s.server.Shutdown()
	return nil
}

func (s *Server) readHandler(filename string, rf io.ReaderFrom) error {
	file, err := s.root.Open(filename)
	if err != nil {
		return err
	}
	_, err = rf.ReadFrom(file)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) writeHandler(filename string, wt io.WriterTo) error {
	return fs.ErrPermission
}
