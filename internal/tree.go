package internal

import (
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/ddddddO/gtree"
	"github.com/fatih/color"
)

// ref: https://github.com/ddddddO/gtree#the-program-below-converts-the-result-of-find-into-a-tree
func tree(bucket string, objList []*storage.ObjectAttrs, option *PrintOption) (string, error) {
	if option.WithColorized {
		bucket = color.BlueString("%s", bucket)
	}
	root := gtree.NewRoot(bucket)
	node := root
	for _, obj := range objList {
		_, file := filepath.Split(obj.Name)
		for _, s := range strings.Split(obj.Name, "/") {
			if s == "" {
				continue
			}
			if option.WithColorized {
				if s == file {
					s = color.GreenString("%s", s)
				} else {
					s = color.BlueString("%s", s)
				}
			}
			tmp := node.Add(s)
			node = tmp
		}
		node = root
	}

	buf := new(strings.Builder)
	if err := gtree.OutputProgrammably(buf, root); err != nil {
		return "", err
	}
	return buf.String(), nil
}
