package tester

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	. "github.com/beanz/everybodycodes/lib-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Run(t *testing.T, fn func([]byte, []byte, []byte) (int, int, int)) {
	d, err := os.ReadFile("TC.txt")
	require.NoError(t, err)
	for _, chunk := range bytes.Split(d, []byte("\n---END---\n")) {
		s := bytes.Split(chunk, []byte("\n"))
		t.Run(string(s[0]), func(t *testing.T) {
			_, p1 := ChompUInt[int](s[1], 0)
			_, p2 := ChompUInt[int](s[2], 0)
			_, p3 := ChompUInt[int](s[3], 0)
			f := string(s[0])
			if f == "ec" {
				f = ""
			}
			i1 := Input(1, f)
			i2 := Input(2, f)
			i3 := Input(3, f)
			a1, a2, a3 := fn(i1, i2, i3)
			assert.Equal(t, p1, a1, "part 1")
			assert.Equal(t, p2, a2, "part 2")
			assert.Equal(t, p3, a3, "part 2")
		})
	}
}

func RunAny[T any, U any, V any](t *testing.T, fn func([]byte, []byte, []byte) (T, U, V)) {
	d, err := os.ReadFile("TC.txt")
	require.NoError(t, err)
	for _, chunk := range bytes.Split(d, []byte("\n---END---\n")) {
		s := bytes.Split(chunk, []byte("\n"))
		t.Run(string(s[0]), func(t *testing.T) {
			f := string(s[0])
			if f == "ec" {
				f = ""
			}
			i1 := Input(1, f)
			i2 := Input(2, f)
			i3 := Input(3, f)
			a1, a2, a3 := fn(i1, i2, i3)
			assert.Equal(t, string(s[1]), fmt.Sprintf("%v", a1), "part 1")
			assert.Equal(t, string(s[2]), fmt.Sprintf("%v", a2), "part 2")
			assert.Equal(t, string(s[3]), fmt.Sprintf("%v", a3), "part 3")
		})
	}
}
