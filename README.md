# 📦 

**Depot** is a versatile command-line tool designed for seamless file synchronization between a server and clients. It stands out by interfacing with file system abstractions, extending beyond conventional OS file structures. For example, you can simply sync the files between your OS file system, Google Drive, and even Telegram.

## Usage

```bash
$ depot server "./videos"
Depot is running on ::1338
```

```bash
$ depot client "::1138" -o "/depot"
Depot client is successfully connected!
```

## Concept of depot/fs

```go
// Package fs contains of Depot filesystem interfaces, where
// each of the subpackages is an implementation of these interfaces.
// It is being used by the Depot's client and server modules.
// The FS interface provides a basic set of functions to interact with
// any kind filesystem-like environment.
package fs

import (
	"io"
	"io/fs"
)

// FS is a read-only fs used primarily by the server.
type FS interface {
	fs.FS

	// Meta returns the fs metadata.
	Meta() (Meta, error)
}

// WriteFS is a read-write fs used primarily by the client.
type WriteFS interface {
	FS

	// WriteFile writes the file using io.WriterTo.
	WriteFile(name string, w io.WriterTo) error
}

// Meta represents useful data of the fs.
type Meta struct {
	// Files is a list of all the files on the fs.
	Files []string
}
```

## Credits

Depot itself is an effort of a mentoring experience led by [@demget](https://github.com/demget). Unfortunately, it stayed unfinished.

The goals of the project were to:
1) Get the very first Go programming experience for my mentee;
2) Get to know with the std package like net, io, bytes, encoding, etc.;
3) Work with several non-std popular packages like cobra, tftp, etc.;
4) Learn how to organize the complex project structure;
5) Learn how to work in a team, using Git, CI, and code review processes;
6) Simply do something interesting and unordinary.
