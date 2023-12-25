package fs

// Server represents a server that shares the FS with its clients.
type Server interface {
	Start() error
	Stop() error
}

// Client represents a client that is able to read from the remote server.
type Client interface {
	// Receive reads the file by its full path and writes to the FS.
	Receive(path string) (File, error)
}

type FS interface {
	Read(path string) (File, error)
	Write(path string, wt io.WriterTo) (File, error)
}

type File interface {}