package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	var a1 int
	for range 10 {
		a1 = p1.round()
	}

	p2 := parse(i2)
	seen2 := make(map[int]int, 1024)
	var a2 int
	for round := 1; ; round++ {
		a := p2.round()
		seen2[a]++
		if seen2[a] == 2024 {
			a2 = round * a
			break
		}
	}

	p3 := parse(i3)
	seen3 := make(map[string]struct{}, 1024)
	a3 := math.MinInt
	for round := 1; ; round++ {
		a := p3.round()
		a3 = max(a3, a)
		k := p3.String()
		if _, ok := seen3[k]; ok {
			break
		}
		seen3[k] = struct{}{}
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	rows [][]int
	pos  int
}

func (p *parsed) round() int {
	n := p.rows[p.pos][0]
	p.rows[p.pos] = p.rows[p.pos][1:]
	p.pos++
	if p.pos == len(p.rows) {
		p.pos = 0
	}
	l := 2 * len(p.rows[p.pos])
	j := (n - 1) % l
	i := min(j, l-j)
	p.rows[p.pos] = slices.Insert(p.rows[p.pos], i, n)

	res, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d", p.rows[0][0], p.rows[1][0], p.rows[2][0], p.rows[3][0]))
	return res
}

func parse(in []byte) *parsed {
	var rows [][]int
	i := 0
	for ; i < len(in); i++ {
		var n int
		i, n = ChompUInt[int](in, i)
		rows = append(rows, []int{n})
		if in[i] == '\n' {
			i++
			break
		}
	}
	j := 0
	for ; i < len(in); i++ {
		var n int
		i, n = ChompUInt[int](in, i)
		rows[j] = append(rows[j], n)
		if i >= len(in) {
			break
		}
		if in[i] == '\n' {
			j = 0
		} else {
			j++
		}
	}
	return &parsed{rows: rows}
}

func (p *parsed) String() string {
	var sb strings.Builder
	var done bool
	i := 0
	for ; !done; i++ {
		done = true
		for _, r := range p.rows {
			if i < len(r) {
				fmt.Fprintf(&sb, "%d ", r[i])
				done = false
			} else {
				fmt.Fprint(&sb, "  ")
			}
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
