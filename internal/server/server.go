package server

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pin/tftp"
)

const DefaultPort = "1338"

type Server struct {
	dataServer *tftp.Server
	path       string
	addr       string
}

func New(addr, path string) Server {
	var s Server
	s.dataServer = tftp.NewServer(s.readHandler, s.writeHandler)
	s.path = path
	s.addr = addr
	return s
}

func (s *Server) Start() error {
	return s.dataServer.ListenAndServe(s.addr)
}

func (s *Server) readHandler(filename string, rf io.ReaderFrom) error {
	err := s.checkDirAccess(s.path + filename)
	if err != nil {
		return err
	}

	file, err := os.Open(s.path + filename)
	if err != nil {
		return err
	}

	_, err = rf.ReadFrom(file)
	if err != nil {
		return err
	}

	s.dataServer.Shutdown()
	return nil
}

func (s *Server) writeHandler(filename string, wt io.WriterTo) error {
	err := s.checkDirAccess(s.path + filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(s.path+filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}

	_, err = wt.WriteTo(file)
	if err != nil {
		return err
	}

	s.dataServer.Shutdown()
	return nil
}

func (s Server) checkDirAccess(filename string) error {
	rootAbs, err := filepath.Abs(s.path)
	if err != nil {
		return err
	}

	fileAbs, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(fileAbs, rootAbs) {
		return os.ErrPermission
	}

	return nil
}
