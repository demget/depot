package tftp

type File struct {
	*os.File
}

func NewFile(f *os.File) *File {
	return &File{File: f}
}

type FS struct {}

func (fs *FS) Write(path string, wt io.WriteTo) (fs.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	_, err = wt.WriteTo(file)
	if err != nil {
		return nil, err
	}

	return NewFile(file), nil
}