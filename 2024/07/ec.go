package main

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, string, int) {
	p1 := parse(i1)
	var a1 []byte
	{
		s := map[byte]int{}
		for _, dev := range p1.devices {
			ds := score(dev.ops, []byte(" "), 10, math.MaxInt)
			s[dev.id] = ds
			a1 = append(a1, dev.id)
		}
		slices.SortFunc(a1, func(a byte, b byte) int {
			return cmp.Compare(s[b], s[a])
		})
	}

	p2 := parse(i2)
	track := parseTrack(testTrack)
	if len(p2.devices) > 4 {
		track = parseTrack(part2Track)
	}
	var a2 []byte
	{
		s := map[byte]int{}
		for _, dev := range p2.devices {
			ds := score(dev.ops, track, 10, math.MaxInt)
			s[dev.id] = ds
			a2 = append(a2, dev.id)
		}
		slices.SortFunc(a2, func(a byte, b byte) int {
			return cmp.Compare(s[b], s[a])
		})
	}

	p3 := parse(i3)
	track = parseTrack(part3Track)
	a3 := 0
	{
		rounds := 11
		target := score(p3.devices[0].ops, track, rounds, math.MaxInt)
		for _, perm := range Permutations([]byte("==="), []byte("---"), []byte("+++++")) {
			s := score(perm, track, rounds, target)
			if s > target {
				a3++
			}
		}
	}

	return string(a1), string(a2), a3
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

func score(ops []byte, track []byte, rounds int, target int) int {
	ds := 0
	power := 10
	i := 0
	for range rounds {
		for _, ch := range track {
			action := ops[i%len(ops)]
			if ch == '+' || ch == '-' {
				action = ch
			}
			switch action {
			case '+':
				power++
			case '-':
				power--
			case '=':
				// no op
			}
			ds += power
			i++
		}
		if ds > target {
			return ds
		}
	}
	return ds
}

type (
	parsed struct {
		devices []*device
	}
	device struct {
		id  byte
		ops []byte
	}
)

func parse(in []byte) *parsed {
	p := &parsed{}
	for line := range bytes.SplitSeq(in, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		id := line[0]
		var ops []byte
		for i := 2; i < len(line); i += 2 {
			ops = append(ops, line[i])
		}
		p.devices = append(p.devices, &device{id, ops})
	}
	return p
}

func (p *parsed) String() string {
	var sb strings.Builder
	for _, dev := range p.devices {
		fmt.Fprintf(&sb, "%c %s\n", dev.id, dev.ops)
	}
	return sb.String()
}

//go:embed test-track.txt
var testTrack []byte

//go:embed part2-track.txt
var part2Track []byte

//go:embed part3-track.txt
var part3Track []byte

func parseTrack(in []byte) []byte {
	w := bytes.IndexByte(in, '\n')
	h := len(in) / (w + 1)
	index := func(x, y int) int {
		return x + (w+1)*y
	}
	get := func(x, y int) byte {
		if !(0 <= x && x < w && 0 <= y && y < h) {
			return ' '
		}
		return in[index(x, y)]
	}
	type dir struct {
		x, y int
	}
	x, y := 1, 0
	d := dir{1, 0}
	var res []byte
	for x != 0 || y != 0 {
		res = append(res, get(x, y))
		nx, ny := x+d.x, y+d.y
		if get(nx, ny) != ' ' {
			x, y = nx, ny
			continue
		}
		d.x, d.y = d.y, d.x
		nx, ny = x+d.x, y+d.y
		if get(nx, ny) != ' ' {
			x, y = nx, ny
			continue
		}
		d.x, d.y = -d.x, -d.y
		nx, ny = x+d.x, y+d.y
		if get(nx, ny) != ' ' {
			x, y = nx, ny
			continue
		}
		panic("bad path")
	}
	res = append(res, 'S')
	return res
}
