package internal

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const GCSTREE_VERSION = "0.0.4"

type GCSTree struct {
	bucket  string
	folder  string
	client  *storage.Client
	counter *counter
}

func NewGCSTree(ctx context.Context, client *storage.Client, bucket string) (*GCSTree, error) {

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

func (g *GCSTree) GetObjectAttrList(ctx context.Context) ([]*storage.ObjectAttrs, error) {
	bkt := g.client.Bucket(g.bucket)
	query := &storage.Query{Prefix: g.folder}
	var names []*storage.ObjectAttrs
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
		names = append(names, attrs)
	}
	return names, nil
}

func (g *GCSTree) String() (string, error) {
	ctx := context.Background()
	objList, err := g.GetObjectAttrList(ctx)
	if err != nil {
		return "", err
	}
	treeResult, err := tree(g.bucket, objList)
	if err != nil {
		return "", err
	}
	treeResult += fmt.Sprintf("\n%s", g.counter.summary())
	return treeResult, nil
}
