package main

import (
	_ "embed"
	"testing"

	. "github.com/beanz/everybodycodes/lib-go"
	"github.com/beanz/everybodycodes/lib-go/tester"
)

func TestParts(t *testing.T) {
	tester.RunAny(t, parts)
}

func BenchmarkParts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		i1 := Input(1, "")
		i2 := Input(2, "")
		i3 := Input(3, "")
		b.StartTimer()
		_, _, _ = parts(i1, i2, i3)
	}
}
