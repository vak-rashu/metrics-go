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

// understand embedded interface
func openPath(path string) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return io.NopCloser(bytes.NewReader(nil)), fmt.Errorf("error opening file %s, Error:\n:%v", path, err)
	}

	// file implements both Reader and Closer interface
	return file, err
}
