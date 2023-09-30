package internal

import (
	"testing"
)

func TestCounter_Summary(t *testing.T) {
	paths := []string{
		"empty_directory/",
		"a.txt",
		"a.tgz",
		"source/a.tgz",
		"source/b.tgz",
		"source/c.tgz",
		"tmp_directory/",
		"tmp_directory/a.png",
		"tmp_directory/empty_directory/",
		"tmp_directory/source/a.tgz",
	}
	want := "5 directories, 7 files"
	c := newCounter()
	for _, p := range paths {
		c.count(p)
	}
	got := c.summary()

	if got != want {
		t.Errorf("\ngot: \n%s\nwant: \n%s", got, want)
	}
}
