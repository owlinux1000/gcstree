package internal

import (
	"fmt"
	"path/filepath"
	"strings"
)

type counter struct {
	dirs  map[string]struct{}
	files map[string]struct{}
}

func newCounter() *counter {
	return &counter{
		dirs:  map[string]struct{}{},
		files: map[string]struct{}{},
	}
}

func (c *counter) count(path string) {
	dir, file := filepath.Split(path)
	if len(dir) > 0 {
		splited := strings.Split(dir, "/")
		for i := range splited {
			c.dirs[filepath.Join(splited[:i+1]...)] = struct{}{}
		}
	}
	if len(file) > 0 {
		c.files[path] = struct{}{}
	}
}

func (c *counter) summary() string {
	return fmt.Sprintf("%d directories, %d files", len(c.dirs), len(c.files))
}
