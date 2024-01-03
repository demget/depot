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
package fs

import (
	"io"
    "io/fs"
)

type FS interface {
    fs.FS
    
    WriteFile(name string, w io.WriterTo) error
}
``` 

## TODO

1. Document all the code inside `fs` and `internal` packages.
2. Test whether the current implementation sync the whole dir.
3. Implement a hidden system file `.depot` to store the metadata (as http communication channel alternative).
