package main

import (
	"bytes"
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

var offsets = []int{0, -1, 0, 1, 0}

func parts(i1, i2, i3 []byte) (int, int, int) {
	return solve(i1, false), solve(i2, false), solve(i3, true)
}

func solve(in []byte, part3 bool) int {
	w := 1 + bytes.IndexByte(in, '\n')
	h := len(in) / (w - 1)
	is := bytes.IndexByte(in, 'S')
	in[is] = '0'
	sx, sy := is%w, is/w
	ie := bytes.IndexByte(in, 'E')
	in[ie] = '0'
	ex, ey := ie%w, ie/w
	type rec struct {
		x, y int
		ch   byte
	}
	inBounds := func(x, y int) bool {
		return 0 <= x && x < w-1 && 0 <= y && y < h
	}
	get := func(x, y int) byte {
		if inBounds(x, y) {
			return in[x+y*w]
		}
		return '#'
	}
	work := make([][]rec, 1024)
	if part3 {
		for x := 1; x < w-2; x++ {
			work[0] = append(work[0], rec{x, 0, 0}, rec{x, h - 1, 0})
			in[x] = '0'
			in[x+(h-1)*w] = '0'
		}
		for y := 1; y < h-1; y++ {
			work[0] = append(work[0], rec{0, y, 0}, rec{w - 2, y, 0})
			in[y*w] = '0'
			in[(w-2)+y*w] = '0'
		}
	} else {
		work[0] = append(work[0], rec{sx, sy, 0})
	}
	for qi := 0; qi < len(work); qi++ {
		for qj := 0; qj < len(work[qi]); qj++ {
			cur := work[qi][qj]
			if cur.x == ex && cur.y == ey {
				return qi
			}
			in[cur.x+cur.y*w] = '#'
			for k := 0; k < 4; k++ {
				nx, ny := cur.x+offsets[k], cur.y+offsets[k+1]
				ch := get(nx, ny)
				if ch == '#' || ch == 'S' {
					continue
				}
				diff := Abs(int(cur.ch) - int(ch-'0'))
				ns := qi + 1 + min(diff, 10-diff)
				work[ns] = append(work[ns], rec{nx, ny, ch - '0'})
			}
		}
	}
	return -1
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
