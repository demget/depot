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

1. Create the proper `fs` package concepts: `FS`, `File`, `Client`, `Server`.
2. Create a `fs/tftp` server implementation.
3. Remeber about the review comments and style in general, document your code.
