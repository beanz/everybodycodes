package main

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, int, int) {
	p1 := parse(i1)
	a1 := string(p1.grids[0].word())
	p2 := parse(i2)
	a2 := p2.power()
	a3 := wall(i3)
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

func wall(in []byte) int {
	w := bytes.IndexByte(in, '\n') + 1
	type todo struct {
		x, y int
	}
	h := len(in) / w
	get := func(x, y int) byte {
		return in[x+w*y]
	}
	set := func(x, y int, ch byte) {
		in[x+w*y] = ch
	}
	progress := true
	for progress {
		progress = false
		for y := 2; y < h; y++ {
			yo := y % 6
			yr := y - yo
			for x := 2; x < w-1; x++ {
				xo := x % 6
				xr := x - xo
				v := []byte{get(x, yr+0), get(x, yr+1), get(x, yr+6), get(x, yr+7)}
				vUnsolved := missing(v)
				h := []byte{get(xr+0, y), get(xr+1, y), get(xr+6, y), get(xr+7, y)}
				hUnsolved := missing(h)
				if get(x, y) == '.' {
					o := filterQuery(intersect(h, v))
					if len(o) == 1 {
						set(x, y, o[0]+32)
						progress = true
					}
				}
				vm := []byte{get(x, yr+2), get(x, yr+3), get(x, yr+4), get(x, yr+5)}
				vmUnsolved := missing(vm)
				hm := []byte{get(xr+2, y), get(xr+3, y), get(xr+4, y), get(xr+5, y)}
				hmUnsolved := missing(hm)
				if get(x, y) == '.' {
					if vUnsolved == 0 && vmUnsolved == 1 {
						o := difference(vm, v)
						if o != 0 {
							set(x, y, o+32)
							progress = true
						}
					} else if hUnsolved == 0 && hmUnsolved == 1 {
						o := difference(hm, h)
						if o != 0 {
							set(x, y, o+32)
							progress = true
						}
					}
				}
				vm = []byte{get(x, yr+2), get(x, yr+3), get(x, yr+4), get(x, yr+5)}
				vmUnsolved = missing(vm)
				if vUnsolved == 1 && vmUnsolved == 0 {
					o := difference(v, vm)
					if o != 0 {
						progress = true
						switch slices.Index(v, '?') {
						case 0:
							set(x, yr+0, o)
						case 1:
							set(x, yr+1, o)
						case 2:
							set(x, yr+6, o)
						case 3:
							set(x, yr+7, o)
						}
					}
				}
				hm = []byte{get(xr+2, y), get(xr+3, y), get(xr+4, y), get(xr+5, y)}
				hmUnsolved = missing(hm)
				if hUnsolved == 1 && hmUnsolved == 0 {
					o := difference(h, hm)
					if o != 0 {
						switch slices.Index(h, '?') {
						case 0:
							set(xr+0, y, o)
						case 1:
							set(xr+1, y, o)
						case 2:
							set(xr+6, y, o)
						case 3:
							set(xr+7, y, o)
						}
						progress = true
					}
				}

				if xo == 5 {
					x += 2
				}
			}
			if yo == 5 {
				y += 2
			}
		}
	}
	res := 0
	for y := 2; y < h; y += 6 {
	WORD:
		for x := 2; x < w-1; x += 6 {
			i := 1
			sum := 0
			var word []byte
			for yy := range 4 {
				for xx := range 4 {
					ch := get(x+xx, y+yy)
					if ch == '.' {
						continue WORD
					}
					word = append(word, ch)
					sum += i * int(1+ch-'a')
					i++
				}
			}
			res += sum
		}
	}
	return res
}

func filterQuery(a []byte) []byte {
	n := 0
	for _, e := range a {
		if e != '?' {
			a[n] = e
			n++
		}
	}
	return a[:n]
}

func missing(a []byte) int {
	c := 0
	for _, e := range a {
		if e == '?' || e == '.' {
			c++
		}
	}
	return c
}

func difference(incomplete, complete []byte) byte {
	var unused []byte
	completeCount := 0
	for i := range complete {
		if complete[i] == '?' || complete[i] == '.' {
			completeCount++
			continue
		}
		ach := complete[i]
		if ach >= 'a' {
			ach -= 32
		}
		found := false
		for j := range incomplete {
			if incomplete[j] == '?' || incomplete[j] == '.' {
				continue
			}
			bch := incomplete[j]
			if bch >= 'a' {
				bch -= 32
			}
			if ach == bch {
				found = true
				break
			}
		}
		if !found {
			unused = append(unused, ach)
		}
	}
	if len(unused) != 1 {
		return 0
	}
	return unused[0]
}

type grid struct {
	h [][]byte
	v [][]byte
}

func (g *grid) word() (res []byte) {
	for y := range 4 {
		for x := range 4 {
			o := intersect(g.h[y], g.v[x])
			res = append(res, o[0])
		}
	}
	return
}

func (g *grid) power() (res int) {
	for i, ch := range g.word() {
		res += (i + 1) * int(1+ch-'A')
	}
	return
}

func (g *grid) String() string {
	var sb strings.Builder
	for k := range 4 {
		fmt.Fprintf(&sb, "H: %s\n", string(g.h[k]))
	}
	for k := range 4 {
		fmt.Fprintf(&sb, "V: %s\n", string(g.v[k]))
	}
	return sb.String()
}

func intersect[T comparable](a []T, b []T) (r []T) {
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				r = append(r, aa)
			}
		}
	}
	return r
}

type parsed struct {
	grids []*grid
}

func (p *parsed) power() (res int) {
	for _, g := range p.grids {
		res += g.power()
	}
	return
}

func parse(in []byte) *parsed {
	p := &parsed{}
	for i := 0; i < len(in); i++ {
		var starts []int
		var w int
	STARTS:
		for j := 0; i+j < len(in); {
			switch in[i+j] {
			case '*':
				starts = append(starts, j)
				j += 8
			case '\n':
				w = j + 1
				break STARTS
			default:
				j++
			}
		}
		for _, start := range starts {
			o := i + start
			g := &grid{h: make([][]byte, 4), v: make([][]byte, 4)}
			for j := range 4 {
				for _, k := range []int{0, 1, 6, 7} {
					g.v[j] = append(g.v[j], in[o+2+j+w*k])
					g.h[j] = append(g.h[j], in[o+(2+j)*w+k])
				}
			}
			p.grids = append(p.grids, g)
		}
		i += w * 8
	}
	return p
}
