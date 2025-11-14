package main

import (
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	return 0, 0, 0
}

func main() {
	p1, p2, p3 := parts(Input(1, "ex"), Input(2, "ex"), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
}

func parse(in []byte) *parsed {
	p := &parsed{}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "todo")
	return sb.String()
}
