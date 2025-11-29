package ec

import (
	"fmt"
	"os"
	"strings"
)

func Dg(argv ...any) {
	if !DEBUG() {
		return
	}
	Lg(argv...)
}

func Df(format string, argv ...any) {
	if !DEBUG() {
		return
	}
	fmt.Fprintf(os.Stderr, format, argv...)
}

func Lg(argv ...any) {
	var sb strings.Builder
	sb.WriteString("%v")
	for i := 1; i < len(argv); i++ {
		sb.WriteString(" %v")
	}
	sb.WriteByte('\n')
	fmt.Fprintf(os.Stderr, sb.String(), argv...)
}

func Lf(format string, argv ...any) {
	fmt.Fprintf(os.Stderr, format, argv...)
}
