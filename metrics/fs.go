package metrics

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const procFileSystem string = "/proc"
const sysFileSystem string = "/sys"

func procPath(p ...string) string {
	return filepath.Join(append([]string{procFileSystem}, p...)...)
}

func sysPath(p ...string) string {
	return filepath.Join(append([]string{sysFileSystem}, p...)...)
}

func openPath(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return io.NopCloser(bytes.NewReader(nil)), fmt.Errorf("error opening file %s, Error:\n:%v", path, err)
	}

	return file, err
}
