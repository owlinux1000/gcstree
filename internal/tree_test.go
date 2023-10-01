package internal

import (
	"testing"

	"cloud.google.com/go/storage"
)

func TestTree(t *testing.T) {
	var objList []*storage.ObjectAttrs
	objList = append(objList, &storage.ObjectAttrs{
		Name: "folder1/",
		Size: 0,
	})
	objList = append(objList, &storage.ObjectAttrs{
		Name: "folder1/folder1-1/hello.txt",
		Size: 6,
	})
	objList = append(objList, &storage.ObjectAttrs{
		Name: "folder2/folder1-2/",
		Size: 6,
	})
	bucket := "test"
	option := PrintOption{
		WithColorized: false,
		WithSize:      false,
	}
	got, err := tree(bucket, objList, &option)
	if err != nil {
		t.Fatal(err)
	}
	want := `test
├── folder1
│   └── folder1-1
│       └── hello.txt
└── folder2
    └── folder1-2
`
	if got != want {
		t.Errorf("\ngot: \n%s\nwant: \n%s", got, want)
	}

}
