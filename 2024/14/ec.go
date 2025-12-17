package main

import (
	"fmt"
	"math"
	"math/bits"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	a1 := 0
	{
		y := 0
		for i := 0; i < len(i1); i++ {
			ch := i1[i]
			j, n := ChompUInt[int](i1, i+1)
			i = j
			switch ch {
			case 'U':
				y += n
				a1 = max(y, a1)
			case 'D':
				y -= n
			}
		}
	}
	var a2 int
	{
		x, y, z := 0, 0, 0
		m := map[[3]int]struct{}{}
		for i := 0; i < len(i2); i++ {
			ch := i2[i]
			j, n := ChompUInt[int](i2, i+1)
			i = j
			switch ch {
			case 'U':
				for i := 0; i < n; i++ {
					y += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'D':
				for i := 0; i < n; i++ {
					y -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'L':
				for i := 0; i < n; i++ {
					x -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'R':
				for i := 0; i < n; i++ {
					x += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'F':
				for i := 0; i < n; i++ {
					z -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'B':
				for i := 0; i < n; i++ {
					z += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			}
			if i < len(i2) && i2[i] == '\n' {
				x, y, z = 0, 0, 0
			}
		}
		a2 = len(m)
	}
	var a3 int
	{
		x, y, z := 0, 0, 0
		m := map[[3]int]struct{}{}
		leaves := map[[3]int]int{}
		mn := [3]int{0, 0, 0}
		mx := [3]int{0, 0, 0}
		for i := 0; i < len(i3); i++ {
			ch := i3[i]
			j, n := ChompUInt[int](i3, i+1)
			i = j
			switch ch {
			case 'U':
				for i := 0; i < n; i++ {
					y += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
				mx[1] = max(y, mx[1])
			case 'D':
				for i := 0; i < n; i++ {
					y -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
			case 'L':
				for i := 0; i < n; i++ {
					x -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
				mn[0] = min(x, mn[0])
			case 'R':
				for i := 0; i < n; i++ {
					x += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
				mx[0] = max(x, mx[0])
			case 'F':
				for i := 0; i < n; i++ {
					z -= 1
					m[[3]int{x, y, z}] = struct{}{}
				}
				mn[2] = min(z, mn[2])
			case 'B':
				for i := 0; i < n; i++ {
					z += 1
					m[[3]int{x, y, z}] = struct{}{}
				}
				mx[2] = max(z, mx[2])
			}
			if i < len(i3) && i3[i] == '\n' {
				leaves[[3]int{x, y, z}] = len(leaves)
				x, y, z = 0, 0, 0
			}
		}
		leaves[[3]int{x, y, z}] = len(leaves)
		a3 = math.MaxInt
		for y := 1; y <= mx[1]; y++ {
			s := dist(m, y, leaves, mn, mx)
			a3 = min(a3, s)
		}
	}
	return a1, a2, a3
}

type Rec struct {
	pos   [3]int
	steps int
}

func dist(m map[[3]int]struct{}, sy int, leaves map[[3]int]int, mn, mx [3]int) int {
	s := 0
	var done uint
	seen := map[[3]int]struct{}{}
	todo := []Rec{{pos: [3]int{0, sy, 0}, steps: 0}}
	for len(todo) > 0 {
		cur := todo[0]
		todo = todo[1:]
		if _, ok := seen[cur.pos]; ok {
			continue
		}
		seen[cur.pos] = struct{}{}
		if n, ok := leaves[cur.pos]; ok {
			bit := uint(1) << n
			if done&bit == 0 {
				s += cur.steps
				done |= bit
				if bits.OnesCount(done) == len(leaves) {
					return s
				}
			}
		}
		add := func(n [3]int) {
			if !(mn[0] <= n[0] && n[0] <= mx[0]) {
				return
			}
			if !(mn[1] <= n[1] && n[1] <= mx[1]) {
				return
			}
			if !(mn[2] <= n[2] && n[2] <= mx[2]) {
				return
			}
			if _, ok := m[n]; !ok {
				return
			}
			todo = append(todo, Rec{pos: n, steps: cur.steps + 1})
		}
		add([3]int{cur.pos[0] - 1, cur.pos[1], cur.pos[2]})
		add([3]int{cur.pos[0] + 1, cur.pos[1], cur.pos[2]})
		add([3]int{cur.pos[0], cur.pos[1] - 1, cur.pos[2]})
		add([3]int{cur.pos[0], cur.pos[1] + 1, cur.pos[2]})
		add([3]int{cur.pos[0], cur.pos[1], cur.pos[2] - 1})
		add([3]int{cur.pos[0], cur.pos[1], cur.pos[2] + 1})
	}
	return math.MaxInt
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}
