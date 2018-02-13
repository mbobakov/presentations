package escape

import (
	"testing"
)

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(click)
		m := make(map[string]*click, 0)
		m["foo"] = c
	}
}
func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(click)
		s := make([]*click, 1)
		s[0] = c
	}
}
