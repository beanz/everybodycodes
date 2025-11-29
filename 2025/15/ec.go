package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := solve(i1)
	a2 := solve(i2)
	a3 := solve(i3)
	return a1, a2, a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type wall struct {
	dx, dy, n int
}

type rec struct {
	x, y, steps int
}

func solve(in []byte) int {
	walls := []wall{}
	xs := map[int]struct{}{-1: {}, 0: {}, 1: {}}
	ys := map[int]struct{}{-1: {}, 0: {}, 1: {}}
	x, y, dx, dy := 0, 0, 0, -1
	for i := 0; i < len(in); i++ {
		if in[i] == 'R' {
			dx, dy = -dy, dx
		} else {
			dx, dy = dy, -dx
		}
		i++
		var n int
		i, n = ChompUInt[int](in, i)
		walls = append(walls, wall{dx, dy, n})
		x += dx * n
		y += dy * n
		xs[x-1] = struct{}{}
		xs[x] = struct{}{}
		xs[x+1] = struct{}{}
		ys[y-1] = struct{}{}
		ys[y] = struct{}{}
		ys[y+1] = struct{}{}
	}

	cx := map[int]int{}
	xk := slices.Collect(maps.Keys(xs))
	sort.Ints(xk)
	for i, x := range xk {
		cx[x] = i
	}
	cy := map[int]int{}
	yk := slices.Collect(maps.Keys(ys))
	sort.Ints(yk)
	for i, y := range yk {
		cy[y] = i
	}
	w := len(cx)
	h := len(cy)
	Dg(w, h)
	Dg(cx, cy)
	m := make([]bool, w*h)
	x, y = 0, 0
	for _, wall := range walls {
		nx, ny := x+wall.n*wall.dx, y+wall.n*wall.dy
		for xc := cx[x]; ; xc += wall.dx {
			for yc := cy[y]; ; yc += wall.dy {
				m[xc+yc*w] = true
				if yc == cy[ny] {
					break
				}
			}
			if xc == cx[nx] {
				break
			}
		}
		x, y = nx, ny
	}
	todo := []rec{{cx[x], cy[y], 0}}
	tx, ty := cx[0], cy[0]
	m[tx+ty*w] = false // remove wall at target
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		Dg(cur)
		if cur.x == tx && cur.y == ty {
			return cur.steps
		}
		for _, o := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			nx, ny := cur.x+o[0], cur.y+o[1]
			if !(0 <= nx && nx < len(xk) && 0 <= ny && ny < len(yk)) {
				continue
			}
			if m[nx+ny*w] {
				continue
			}
			m[nx+ny*w] = true
			steps := cur.steps + Abs(xk[cur.x]-xk[nx]) + Abs(yk[cur.y]-yk[ny])
			todo = append(todo, rec{nx, ny, steps})
		}
	}

	return 0
}

func dump(m []bool, w, h int, todo []rec) string {
	var sb strings.Builder
	for y := range h {
		for x := range w {
			if slices.ContainsFunc(todo, func(e rec) bool {
				return e.x == x && e.y == y
			}) {
				fmt.Fprintf(&sb, "+")
			} else if m[x+y*w] {
				fmt.Fprintf(&sb, "#")
			} else {
				fmt.Fprintf(&sb, ".")
			}
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
