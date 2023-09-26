package internal

import (
	"context"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/ddddddO/gtree"
	"google.golang.org/api/iterator"
)

type GCSTree struct {
	bucket string
	client *storage.Client
}

func NewGCSTree(ctx context.Context, bucket string) (*GCSTree, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GCSTree{
		bucket: bucket,
		client: client,
	}, nil
}

func (g *GCSTree) GetObjectList(ctx context.Context) ([]string, error) {

	bkt := g.client.Bucket(g.bucket)
	query := &storage.Query{Prefix: ""}
	var names []string
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		names = append(names, attrs.Name)
	}
	return names, nil
}

// ref: https://github.com/ddddddO/gtree#the-program-below-converts-the-result-of-find-into-a-tree
func (g *GCSTree) Tree() (string, error) {
	ctx := context.Background()
	objList, err := g.GetObjectList(ctx)
	if err != nil {
		return "", err
	}
	root := gtree.NewRoot(g.bucket)
	node := root
	for _, obj := range objList {
		for _, s := range strings.Split(obj, "/") {
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
