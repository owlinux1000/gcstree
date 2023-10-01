package internal

import "testing"

func TestFormatBytes(t *testing.T) {
	tests := map[int64]string{
		999:       "999",
		1024:      "1.0K",
		1536:      "1.5K",
		987654321: "987.7M",
	}
	for key, value := range tests {
		got := formatBytes(key)
		want := value
		if got != want {
			t.Fatalf("got: %v, want: %v\n", got, want)
		}
	}
}
