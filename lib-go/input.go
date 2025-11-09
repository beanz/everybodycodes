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
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	top, year, day := findDir(os.Args[0], cwd)
	inputPath := filepath.Join(top, "input", year, day)
	if ex != "" {
		return filepath.Join(inputPath, fmt.Sprintf("%s-p%d.txt", ex, n))
	}
	return filepath.Join(inputPath, fmt.Sprintf("everybody_codes_e%s_q%s_p%d.txt", year, day, n))
}

func findDir(exe, cwd string) (string, string, string) {
	abs, err := filepath.Abs(exe)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(abs)
	day := filepath.Base(dir)
	year := filepath.Base(filepath.Dir(dir))
	if len(day) == 2 && len(year) == 4 {
		return filepath.Dir(filepath.Dir(dir)), year, day
	}
	day = filepath.Base(cwd)
	year = filepath.Base(filepath.Dir(cwd))
	return filepath.Dir(filepath.Dir(cwd)), year, day
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

func ChompInt[T Signed](in []byte, i int) (int, T) {
	j, n := i, T(0)
	var negative bool
	if in[j] == '-' {
		negative = true
		j++
	} else if in[j] == '+' {
		j++
	}
	if !('0' <= in[j] && in[j] <= '9') {
		panic("not a number")
	}
	for ; j < len(in) && '0' <= in[j] && in[j] <= '9'; j++ {
		n = T(10)*n + T(in[j]&0xf)
	}
	if negative {
		return j, -n
	}
	return j, n
}
