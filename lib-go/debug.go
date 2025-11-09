package ec

import (
	"fmt"
	"os"
	"strings"
)

func LG(argv ...any) {
	var sb strings.Builder
	sb.WriteString("%v")
	for i := 1; i < len(argv); i++ {
		sb.WriteString(" %v")
	}
	sb.WriteByte('\n')
	fmt.Fprintf(os.Stderr, sb.String(), argv...)
}
