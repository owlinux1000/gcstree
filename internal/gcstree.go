package internal

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/ddddddO/gtree"
	"google.golang.org/api/iterator"
)

const GCSTREE_VERSION = "0.0.4"

type GCSTree struct {
	bucket  string
	folder  string
	client  *storage.Client
	counter *counter
}

func NewGCSTree(ctx context.Context, bucket string) (*GCSTree, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	folder := ""
	if strings.Contains(bucket, "/") {
		splited := strings.Split(bucket, "/")
		folder = strings.Join(splited[1:], "/")
		bucket = splited[0]
	}
	return &GCSTree{
		bucket:  bucket,
		folder:  folder,
		client:  client,
		counter: newCounter(),
	}, nil
}

func (g *GCSTree) GetObjectList(ctx context.Context) ([]string, error) {
	bkt := g.client.Bucket(g.bucket)
	query := &storage.Query{Prefix: g.folder}
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
		g.counter.count(attrs.Name)
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
	if _, err := buf.WriteString(fmt.Sprintf("\n%s", g.counter.summary())); err != nil {
		return "", err
	}
	return buf.String(), nil
}
