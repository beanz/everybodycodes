package main

import (
	"bytes"
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := part1(i1)
	p2 := part2(i2)
	p3 := part3(i3)
	return p1, p2, p3
}

func main() {
	p1, p2, p3 := parts(Input(1, "ex"), Input(2, ""), Input(3, "ex"))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

func part3(in []byte) int {
	l := len(in)
	if in[l-1] == '\n' {
		l--
	}
	dna := bytes.Split(in[:l], []byte{'\n'})
	res := 0
	f := make([]int, len(dna))
	scores := make([]int, len(dna))
	nextFamily := 1
	join := func(i, j int) {
		fi := f[i]
		fj := f[j]
		if fi == fj {
			return
		}
		family := nextFamily
		nextFamily++
		for k := range len(dna) {
			if f[k] == fi || f[k] == fj {
				f[k] = family
			}
		}
	}
	for i := range len(dna) {
		child := dna[i][bytes.Index(dna[i], []byte{':'})+1:]
		for j := range len(dna) {
			if i == j {
				continue
			}
			p1 := dna[j][bytes.Index(dna[j], []byte{':'})+1:]
			for k := j + 1; k < len(dna); k++ {
				if i == k {
					continue
				}
				p2 := dna[k][bytes.Index(dna[k], []byte{':'})+1:]
				s := score(child, p1, p2)
				if s == 0 {
					continue
				}
				res += s
				scores[i] = s
				childHasFamily := f[i] != 0
				p1HasFamily := f[j] != 0
				p2HasFamily := f[k] != 0
				if !childHasFamily && !p1HasFamily && !p2HasFamily {
					family := nextFamily
					nextFamily++
					f[i] = family
					f[j] = family
					f[k] = family
					continue
				}
				if !childHasFamily {
					family := nextFamily
					nextFamily++
					f[i] = family
				}
				if !p1HasFamily {
					family := nextFamily
					nextFamily++
					f[j] = family
				}
				if !p2HasFamily {
					family := nextFamily
					nextFamily++
					f[k] = family
				}
				join(i, j)
				join(i, k)
			}
		}
	}
	m := make(map[int]int)
	sc := make(map[int]int)
	mx := 0
	mf := 0
	for i := range len(dna) {
		f := f[i]
		if f == 0 {
			continue
		}
		m[f]++
		sc[f] += i + 1
		if m[f] > mx {
			mx = m[f]
			mf = f
		}
	}
	return sc[mf]
}

func part2(in []byte) int {
	l := len(in)
	if in[l-1] == '\n' {
		l--
	}
	dna := bytes.Split(in[:l], []byte{'\n'})
	res := 0
	for i := range len(dna) {
		child := dna[i][bytes.Index(dna[i], []byte{':'})+1:]
		for j := range len(dna) {
			if i == j {
				continue
			}
			p1 := dna[j][bytes.Index(dna[j], []byte{':'})+1:]
			for k := j + 1; k < len(dna); k++ {
				if i == k {
					continue
				}
				p2 := dna[k][bytes.Index(dna[k], []byte{':'})+1:]
				s := score(child, p1, p2)
				res += s
			}
		}
	}
	return res
}

func part1(in []byte) int {
	w := bytes.IndexByte(in, '\n')
	a := in[2:w]
	b := in[w+3 : w+w+1]
	c := in[w+w+4 : w+w+w+2]
	a1 := score(a, b, c)
	if a1 != 0 {
		return a1
	}
	a1 = score(b, a, c)
	if a1 != 0 {
		return a1
	}
	return score(c, a, b)
}

func score(child, p1, p2 []byte) int {
	var s1, s2 int
	for i := range len(child) {
		var match bool
		if child[i] == p1[i] {
			s1++
			match = true
		}
		if child[i] == p2[i] {
			s2++
			match = true
		}
		if !match {
			return 0
		}
	}
	return s1 * s2
}
