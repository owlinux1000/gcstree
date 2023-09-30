package internal

import (
	"testing"
)

func TestCounter_Summary(t *testing.T) {
	tests := []struct {
		name  string
		paths []string
		want  string
	}{
		{
			name: "pattern1",
			paths: []string{
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
			},
			want: "5 directories, 7 files",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := newCounter()
			for _, p := range tt.paths {
				c.count(p)
			}
			got := c.summary()

			if got != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.want)
			}
		})
	}
}
