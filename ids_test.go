package ids_test

import (
	"testing"

	"github.com/dotpy3/ids"
)

const (
	v3ID = "v3/gtw"
	v2ID = "gtw"
)

func BenchmarkReadIDUsingStrings(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, v3 := ids.ReadIDUsingStrings(v2ID)
		if v3 {
			b.FailNow()
		}
	}
	for n := 0; n < b.N; n++ {
		_, v3 := ids.ReadIDUsingStrings(v3ID)
		if !v3 {
			b.FailNow()
		}
	}
}

func BenchmarkReadIDUsingRegexp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, v3 := ids.ReadIDUsingRegexp(v2ID)
		if v3 {
			b.FailNow()
		}
	}
	for n := 0; n < b.N; n++ {
		_, v3 := ids.ReadIDUsingRegexp(v3ID)
		if !v3 {
			b.FailNow()
		}
	}
}
