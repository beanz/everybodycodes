package main

import (
	"bytes"
	"fmt"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.count(10)
	p2 := parse(i2)
	var d int
	var dr int
	for r := range min(p2.w/2, p2.h/2) {
		c := p2.count(r + 1)
		if d < c {
			d = c
			dr = r + 1
		}
	}
	p3 := parse(i3)
	a3 := p3.solve()
	return a1, d * dr, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	in     []byte
	w, h   int
	ox, oy int
	sx, sy int
}

func parse(in []byte) *parsed {
	p := &parsed{in: in}
	p.w = 1 + bytes.IndexByte(in, '\n')
	p.h = (1 + len(in)) / p.w
	i := bytes.IndexByte(in, '@')
	p.ox, p.oy = i%p.w, i/p.w
	i = bytes.IndexByte(in, 'S')
	p.sx, p.sy = i%p.w, i/p.w
	return p
}

func (p *parsed) rd(x, y int) int {
	return (p.ox-x)*(p.ox-x) + (p.oy-y)*(p.oy-y)
}

func (p *parsed) rr(t int) int {
	r := int(t / 30)
	return r * r
}

func (p *parsed) winding(x, _, nx, ny, w int) int {
	if ny <= p.oy {
		return w
	}
	if x <= p.ox && nx > p.ox {
		return w + 1
	}
	if x > p.ox && nx <= p.ox {
		return w - 1
	}
	return w
}

func (p *parsed) solve() int {
	type rec struct {
		x, y    int
		t       int
		z       int
		closest int
	}
	work := [205209][]rec{0: {{p.sx, p.sy, 0, 0, p.rd(p.sx, p.sy)}}}
	seen := map[[3]int]int{}
	for qi := 0; qi < len(work); qi++ {
		for qj := 0; qj < len(work[qi]); qj++ {
			cur := work[qi][qj]
			if cur.x == p.sx && cur.y == p.sy && cur.z == 1 {
				return cur.t * int(cur.t/30)
			}
			k := [3]int{cur.x, cur.y, cur.z}
			if closest, ok := seen[k]; ok && closest >= cur.closest {
				continue
			}
			seen[k] = cur.closest
			for _, o := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				nx, ny := cur.x+o[0], cur.y+o[1]
				if !(0 <= nx && nx < p.w-1 && 0 <= ny && ny < p.h) {
					continue
				}
				if nx == p.ox && ny == p.oy {
					continue
				}
				var nt int
				if nx == p.sx && ny == p.sy {
					nt = cur.t
				} else {
					nt = cur.t + int(p.in[nx+ny*p.w]-'0')
				}
				rd := p.rd(nx, ny)
				nc := min(cur.closest, rd)
				if nc <= p.rr(nt) {
					continue
				}
				nz := p.winding(cur.x, cur.y, nx, ny, cur.z)
				work[nt] = append(work[nt], rec{nx, ny, nt, nz, nc})
			}
		}
		Dg(qi)
		work[qi] = work[qi][:0]
	}
	return -1
}

func (p *parsed) count(r int) int {
	rr := r * r
	a := 0
	for y := range p.h {
		for x := range p.w - 1 {
			i := x + y*p.w
			ch := p.in[i]
			if ch == '@' {
				continue
			}
			if (x-p.ox)*(x-p.ox)+(y-p.oy)*(y-p.oy) <= rr {
				a += int(ch - '0')
				p.in[i] = '@'
			}
		}
	}
	return a
}

func (p *parsed) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%dx%d %d,%d %d,%d\n", p.w, p.h, p.ox, p.oy, p.sx, p.sy)
	fmt.Fprintln(&sb, string(p.in))
	return sb.String()
}
