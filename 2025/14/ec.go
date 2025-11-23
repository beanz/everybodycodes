package main

import (
	"bytes"
	"fmt"
	"math/bits"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := parse(i1).rounds(10)
	a2 := parse(i2).rounds(2025)
	m3 := parse(i3)
	m3.shift()
	p3 := &parsed{rows: make([]uint64, 36), w: 34, h: 34}
	a3 := 0
	for range 4095 {
		c := p3.round()
		if p3.matches(m3) {
			a3 += c
		}
	}
	a3 *= (1000000000 / 4095)
	for range 1 + (1000000000 % 4095) {
		c := p3.round()
		if p3.matches(m3) {
			a3 += c
		}
	}
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	rows []uint64
	w, h int
}

func parse(in []byte) *parsed {
	w := bytes.IndexByte(in, '\n') + 1
	h := (1 + len(in)) / w
	p := &parsed{w: w - 1, h: h, rows: []uint64{0}}
	for y := range h {
		var row uint64
		bit := uint64(2)
		for x := range w - 1 {
			if in[x+y*w] == '#' {
				row |= bit
			}
			bit <<= 1
		}
		p.rows = append(p.rows, row)
	}
	p.rows = append(p.rows, 0)
	return p
}

func (p *parsed) shift() {
	p.w += 13
	for y := range p.h {
		p.rows[y+1] <<= 13
	}
}

func (p *parsed) matches(m *parsed) bool {
	mask := uint64(0xff << 14)
	o := 13
	for y := range m.h {
		if (p.rows[o+y+1] & mask) != m.rows[y+1] {
			return false
		}
	}
	return true
}

func (p *parsed) rounds(n int) (c int) {
	for range n {
		c += p.round()
	}
	return
}

func (p *parsed) round() int {
	active := 0
	next := make([]uint64, p.h+2)
	for y := range p.h {
		bitPrev := uint64(1)
		bit := uint64(2)
		bitAfter := uint64(4)
		for range p.w {
			c := 0
			if p.rows[y]&bitPrev != 0 {
				c++
			}
			if p.rows[y]&bitAfter != 0 {
				c++
			}
			if p.rows[y+2]&bitPrev != 0 {
				c++
			}
			if p.rows[y+2]&bitAfter != 0 {
				c++
			}
			active := p.rows[y+1]&bit != 0
			if active {
				if c%2 == 1 {
					next[y+1] |= bit
				}
			} else {
				if c%2 == 0 {
					next[y+1] |= bit
				}
			}
			bit <<= 1
			bitPrev <<= 1
			bitAfter <<= 1
		}
		active += bits.OnesCount(uint(next[y+1]))
	}
	p.rows = next
	return active
}

func (p *parsed) String() string {
	var sb strings.Builder
	f := fmt.Sprintf("%%0%db\n", p.w+2)
	for _, row := range p.rows {
		fmt.Fprintf(&sb, f, row)
	}
	return sb.String()
}
