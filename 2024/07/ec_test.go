package main

import (
	_ "embed"
	"testing"

	. "github.com/beanz/everybodycodes/lib-go"
	"github.com/beanz/everybodycodes/lib-go/tester"

	"github.com/stretchr/testify/assert"
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

func Test_parseTrack(t *testing.T) {
	assert.Equal(t, []byte("+===++-=+=-S"), parseTrack(testTrack))
	expectPart2 := []byte(
		"-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--" +
			"-=++==-" +
			"-+++==++=+=--==++==+++=++=+++=--=+=-=+=-+=-+=-+-=+=-=+=-+++=+==++++==--" +
			"-=+=+=-S")
	assert.Equal(t, expectPart2, parseTrack(part2Track))
	expectPart3 := []byte(`+=+++===-+++++=-==+--+=+===-++=====+--===++=-==+=++====-==-===+=+=--==++=+========-=======++--+++=-++=-+=+==-=++=--+=-====++--+=-==++======+=++=-+==+=-==++=-=-=---++=-=++==++===--==+===++===---+++==++=+=-=====+==++===--==-==+++==+++=++=+===--==++--===+=====-=++====-+=-+--=+++=-+-===++====+++--=++====+=-=+===+=====-+++=+==++++==----=+=+=-S`)
	assert.Equal(t, expectPart3, parseTrack(part3Track))
}
