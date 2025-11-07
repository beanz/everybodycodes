package ec

import (
	"fmt"
	"os"
	"path/filepath"
)

func Input(n int, ex string) []byte {
	bytes, err := os.ReadFile(InputFile(n, ex))
	if err != nil {
		panic(err)
	}
	return bytes
}

func InputFile(n int, ex string) string {
	abs, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(abs)
	day := filepath.Base(dir)
	year := filepath.Base(filepath.Dir(dir))
	top := filepath.Dir(filepath.Dir(dir))
	inputPath := filepath.Join(top, "input", year, day)
	if ex != "" {
		return filepath.Join(inputPath, fmt.Sprintf("%s-p%d.txt", ex, n))
	}
	return filepath.Join(inputPath, fmt.Sprintf("everybody_codes_e%s_q%s_p%d.txt", year, day, n))
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Int interface{ Unsigned | Signed }

func ChompUInt[T Int](in []byte, i int) (j int, n T) {
	j = i
	if !('0' <= in[j] && in[j] <= '9') {
		panic("not a number")
	}
	for ; j < len(in) && '0' <= in[j] && in[j] <= '9'; j++ {
		n = T(10)*n + T(in[j]&0xf)
	}
	return
}
