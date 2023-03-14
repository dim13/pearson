package pearson

import (
	"fmt"
	"strings"
	"testing"
)

func name(t *testing.T, s string) string {
	t.Helper()
	if len(s) > 20 {
		return s[:20] + "..."
	}
	return s
}

func TestHash(t *testing.T) {
	testCases := []struct {
		input  string
		result byte
	}{
		{"", 0},
		{"test", 47},
		{"whatever", 27},
		{"RFC 3074", 214},
		{strings.Repeat("A", 300), 139},
	}
	for _, tc := range testCases {
		t.Run(name(t, tc.input), func(t *testing.T) {
			h := New()
			fmt.Fprint(h, tc.input)
			if s := h.Sum(nil); s[0] != tc.result {
				t.Errorf("want %v, got %v", tc.result, s[0])
			}
		})
	}
}

func BenchmarkHash(b *testing.B) {
	h := New()
	v := []byte("test")
	for i := 0; i < b.N; i++ {
		h.Write(v)
	}
}
