# 📦 Depot

## Usage

```bash
$ depot server "./videos"
Depot is running on ::1338
```

```bash
$ depot client "::1138" -o "/depot"
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

1. Once you start a server with a path to share, the server is waiting for the clients to connect and receive the contents of the specified path immediately.

```bash
$ depot server "./depot/server"
Depot is running on ::1338
```

2. Once a client is connected, the files from the server should start downloading immediately to the specified `-o` output path.

```bash
$ depot client -o "./depot/client" "::1138"
Depot client is successfully connected!
```

3. To solve this problem we need to decide on the best protocol for such data sharing.
4. Put all the internal server and client logic into the `internal` package in the root (`internal/server` and `internal/client`).
