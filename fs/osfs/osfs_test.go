package osfs

import (
	"github.com/demget/depot/fs"
)

var _ = *FS(nil).(fs.FS)