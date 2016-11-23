package pearson

import (
	"fmt"
	"strings"
	"testing"
)

var testCases = []struct {
	input  string
	result byte
}{
	{"", 0},
	{"test", 47},
	{"whatever", 27},
	{"RFC 3074", 214},
	{strings.Repeat("A", 300), 139},
}

func TestHash(t *testing.T) {
	h := New()
	for _, tc := range testCases {
		fmt.Fprint(h, tc.input)
		if s := h.Sum(nil); s[0] != tc.result {
			t.Errorf("expected %v, got %v", tc.result, s[0])
		}
	}
}

func BenchmarkHash(b *testing.B) {
	h := New()
	for i := 0; i < b.N; i++ {
		h.Write([]byte("test"))
	}
}
