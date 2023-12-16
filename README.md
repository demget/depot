# 📦 Depot

## Usage

```bash
$ depot server "./videos"
Depot is running on ::1338
```

```bash
$ depot client "::1138" --path "/root"
Depot client is successfully connected!
```

## Basic concept

```go
type WalkFn = func(string, *File, error) error

type FS interface {
    Open(string) (*File, error)
    Walk(string, WalkFn) error
}

type File interface {
    // ...
}
```

## TODO

Learn about https://github.com/spf13/cobra, try to implement a basic command application, which parses the command, arguments, flags, etc.

```bash
$ depot server "./videos" --addr ":80"
Command: server
Args: ./videos
Flags: addr=:80

$ depot client ":80"
Command: client
Args: :80
Flags:
```
