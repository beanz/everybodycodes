package main

import (
	"bytes"
	"fmt"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.solve1(100)
	p2 := parse(i2)
	a2 := p2.solve2()
	p3 := parse(i3)
	a3 := p3.solve3()
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type parsed struct {
	m    []byte
	w, h int
	sx   int
}

func parse(in []byte) *parsed {
	p := &parsed{}
	p.w = bytes.IndexByte(in, '\n') + 1
	p.h = (1 + len(in)) / p.w
	p.m = in
	p.sx = bytes.IndexByte(in, 'S')
	return p
}

func (p *parsed) solve3() int {
	type Rec struct {
		x, y, z int
		dx, dy  int
		n       int
	}
	todo := []Rec{{p.sx, 0, 100, 0, 1, 0}}
	seen := map[[4]int]int{}
	var mx, my, mz int
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		if cur.z <= 0 {
			if cur.y > my {
				my = cur.y
				mx = cur.x
			}
			continue
		}
		k := [4]int{cur.x, cur.y, cur.dx, cur.dy}
		if z, ok := seen[k]; ok && z >= cur.z {
			continue
		}
		seen[k] = cur.z
		add := func(dx, dy int) {
			x, y, z := cur.x+dx, cur.y+dy, cur.z
			if y < 0 {
				return
			}
			ch := p.m[x+(y%p.h)*p.w]
			//Lf("%d,%d,%d %c\n", x, y, z, ch)
			switch ch {
			case '#':
				return
			case '.', 'S':
				z--
			case '-':
				z -= 2
			case '+':
				z++
			default:
				panic("foo")
			}
			todo = append(todo, Rec{x, y, z, dx, dy, cur.n + 1})
		}
		add(-cur.dy, cur.dx)
		add(cur.dx, cur.dy)
		add(cur.dy, -cur.dx)
	}
	Dg(mx, mz)
	dz := 0
	for y := 0; y < p.h; y++ {
		ch := p.m[mx+y*p.w]
		//Lf("%d,%d,%d %c\n", x, y, z, ch)
		switch ch {
		case '.', 'S':
			dz--
		case '-':
			dz -= 2
		case '+':
			dz++
		default:
			panic("foo")
		}
	}
	Dg(dz)
	target := 384400
	n := target / -dz
	n -= 10
	target += n * dz
	todo = []Rec{{p.sx, 0, target, 0, 1, 0}}
	seen = map[[4]int]int{}
	my = 0
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		if cur.z <= 0 {
			if cur.y > my {
				my = cur.y
			}
			continue
		}
		k := [4]int{cur.x, cur.y, cur.dx, cur.dy}
		if z, ok := seen[k]; ok && z >= cur.z {
			continue
		}
		seen[k] = cur.z
		add := func(dx, dy int) {
			x, y, z := cur.x+dx, cur.y+dy, cur.z
			if y < 0 {
				return
			}
			ch := p.m[x+(y%p.h)*p.w]
			//Lf("%d,%d,%d %c\n", x, y, z, ch)
			switch ch {
			case '#':
				return
			case '.', 'S':
				z--
			case '-':
				z -= 2
			case '+':
				z++
			default:
				panic("foo")
			}
			todo = append(todo, Rec{x, y, z, dx, dy, cur.n + 1})
		}
		add(-cur.dy, cur.dx)
		add(cur.dx, cur.dy)
		add(cur.dy, -cur.dx)
	}

	return my + n*p.h
}

func (p *parsed) solve2() int {
	type Rec struct {
		x, y, z int
		dx, dy  int
		n       int
		f       int
	}
	todo := []Rec{{p.sx, 0, 10000, 0, 1, 0, 0}}
	seen := map[[5]int]int{}
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		//Lg(cur.x, cur.y, cur.z, cur.dx, cur.dy, cur.n)
		if cur.x == p.sx && cur.y == 0 && cur.z >= 10000 && cur.f == 7 {
			return cur.n
		}
		k := [5]int{cur.x, cur.y, cur.dx, cur.dy, cur.f}
		if z, ok := seen[k]; ok && z >= cur.z {
			continue
		}
		seen[k] = cur.z
		add := func(dx, dy int) {
			x, y, z := cur.x+dx, cur.y+dy, cur.z
			if y >= p.h || y < 0 {
				return
			}
			f := cur.f
			ch := p.m[x+y*p.w]
			switch ch {
			case '#':
				return
			case '.', 'S':
				z--
			case 'A':
				z--
				f |= 1
			case 'B':
				z--
				if f == 1 {
					f |= 2
				}
			case 'C':
				z--
				if f == 3 {
					f |= 4
				}
			case '-':
				z -= 2
			case '+':
				z++
			default:
				panic("foo")
			}
			todo = append(todo, Rec{x, y, z, dx, dy, cur.n + 1, f})
		}
		add(-cur.dy, cur.dx)
		add(cur.dx, cur.dy)
		add(cur.dy, -cur.dx)
	}
	return -1
}

func (p *parsed) solve1(n int) int {
	mx := 0
	type Rec struct {
		x, y, z int
		dx, dy  int
		n       int
	}
	todo := []Rec{{p.sx, 0, 1000, 0, 1, 0}}
	seen := map[[4]int]int{}
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		//Lg(cur.x, cur.y, cur.z, cur.dx, cur.dy, cur.n)
		if cur.n == n {
			mx = max(mx, cur.z)
			continue
		}
		k := [4]int{cur.x, cur.y, cur.dx, cur.dy}
		if z, ok := seen[k]; ok && z >= cur.z {
			continue
		}
		seen[k] = cur.z
		add := func(dx, dy int) {
			x, y, z := cur.x+dx, cur.y+dy, cur.z
			if y >= p.h || y < 0 {
				return
			}
			ch := p.m[x+y*p.w]
			switch ch {
			case '#':
				return
			case '.', 'S':
				z--
			case '-':
				z -= 2
			case '+':
				z++
			default:
				panic("foo")
			}
			todo = append(todo, Rec{x, y, z, dx, dy, cur.n + 1})
		}
		add(-cur.dy, cur.dx)
		add(cur.dx, cur.dy)
		add(cur.dy, -cur.dx)
	}
	return mx
}
