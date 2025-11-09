package main

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := Parse(i1)
	a1 := 0
	for _, w := range p1.words {
		a1 += bytes.Count(p1.lines[0], w)
	}
	p2 := Parse(i2)
	a2 := 0
	for _, w := range p2.words {
		rw := make([]byte, len(w))
		wl := len(w) - 1
		for i := range len(w) {
			rw[i] = w[wl-i]
		}
		p2.words = append(p2.words, rw)
	}
	for _, l := range p2.lines {
		for i := range len(l) {
		WORDS:
			for _, w := range p2.words {
				for j := range len(w) {
					if i+j >= len(l) {
						continue WORDS
					}
					if w[j]&0x1f != l[i+j]&0x1f {
						continue WORDS
					}
				}
				for j := range len(w) {
					l[i+j] = 96 + l[i+j]&0x1f
				}
			}
		}
		c := 0
		for i := range len(l) {
			if l[i] > 96 {
				c++
			}
		}
		a2 += c
	}

	p3 := Parse(i3)
	h := len(p3.lines)
	w := len(p3.lines[0])
	for _, w := range p3.words {
		rw := make([]byte, len(w))
		wl := len(w) - 1
		for i := range len(w) {
			rw[i] = w[wl-i]
		}
		p3.words = append(p3.words, rw)
	}
	for y := range h {
		for x := range w {
			for _, word := range p3.words {
				c := 0
				for i := range len(word) {
					nx := (x + i) % w
					if word[i]&0x1f == p3.lines[y][nx]&0x1f {
						c++
					}
				}
				if c == len(word) {
					for i := range len(word) {
						nx := (x + i) % w
						p3.lines[y][nx] = 96 + p3.lines[y][nx]&0x1f
					}
				}
				c = 0
				for i := range len(word) {
					ny := (y + i)
					if ny >= h {
						break
					}
					if word[i]&0x1f == p3.lines[ny][x]&0x1f {
						c++
					}
				}
				if c == len(word) {
					for i := range len(word) {
						ny := (y + i)
						p3.lines[ny][x] = 96 + p3.lines[ny][x]&0x1f
					}
				}
			}
		}
	}
	a3 := 0
	for y := range h {
		for x := range w {
			if p3.lines[y][x] > 96 {
				a3++
			}
		}
	}

	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	words [][]byte
	lines [][]byte
}

func Parse(in []byte) *parsed {
	w, lines, _ := bytes.Cut(in, []byte("\n\n"))
	words := bytes.Split(w[6:], []byte(","))
	return &parsed{
		words: words,
		lines: bytes.Split(bytes.Trim(lines, "\n"), []byte("\n")),
	}
}

func (p *parsed) String() string {
	var sb strings.Builder
	for _, w := range p.words {
		fmt.Fprintf(&sb, "%s,", w)
	}
	fmt.Fprintln(&sb)
	for _, l := range p.lines {
		fmt.Fprintf(&sb, "L: %s\n", l)
	}
	return sb.String()
}
