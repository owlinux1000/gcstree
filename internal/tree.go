package internal

import (
	"strings"

	"cloud.google.com/go/storage"
	"github.com/ddddddO/gtree"
)

// ref: https://github.com/ddddddO/gtree#the-program-below-converts-the-result-of-find-into-a-tree
func tree(bucket string, objList []*storage.ObjectAttrs) (string, error) {
	root := gtree.NewRoot(bucket)
	node := root
	for _, obj := range objList {
		for _, s := range strings.Split(obj.Name, "/") {
			if s == "" {
				continue
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
