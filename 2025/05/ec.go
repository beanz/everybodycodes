package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (int, int, int) {
	p1 := parse(i1)
	a1 := p1.swords[0].quality

	p2 := parse(i2)
	sort.Slice(p2.swords, func(i, j int) bool {
		return p2.swords[i].quality < p2.swords[j].quality
	})
	a2 := p2.swords[len(p2.swords)-1].quality - p2.swords[0].quality

	p3 := parse(i3)
	slices.SortFunc(p3.swords, compare)
	return a1, a2, p3.checksum()
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type (
	parsed struct {
		swords []*sword
	}
	sword struct {
		id, quality int
		fishbone    []*spine
	}
	spine struct {
		left  *int
		mid   int
		right *int
	}
)

func newSword(in []byte, i *int) *sword {
	var id int
	*i, id = ChompUInt[int](in, *i)
	var fb []*spine
OUTER:
	for *i < len(in) && in[*i] != '\n' {
		j, n := ChompUInt[int](in, *i+1)
		*i = j
		for _, s := range fb {
			if s.left == nil && n < s.mid {
				s.left = &n
				continue OUTER
			}
			if s.right == nil && n > s.mid {
				s.right = &n
				continue OUTER
			}
		}
		fb = append(fb, &spine{nil, n, nil})
	}
	var sb strings.Builder
	for j := 0; j < len(fb); j++ {
		fmt.Fprint(&sb, fb[j].mid)
	}
	q, _ := strconv.Atoi(sb.String())
	return &sword{
		id:       id,
		quality:  q,
		fishbone: fb,
	}
}

func compare(b, a *sword) int {
	if n := cmp.Compare(a.quality, b.quality); n != 0 {
		return n
	}
	l := min(len(a.fishbone), len(b.fishbone))
	for i := 0; i < l; i++ {
		if n := cmp.Compare(a.fishbone[i].num(), b.fishbone[i].num()); n != 0 {
			return n
		}
	}
	return cmp.Compare(a.id, b.id)
}

func (s *sword) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%d:\n", s.id)
	for _, spine := range s.fishbone {
		if spine.left == nil {
			sb.WriteString("  ")
		} else {
			fmt.Fprintf(&sb, "%d-", *spine.left)
		}
		fmt.Fprint(&sb, spine.mid)
		if spine.right != nil {
			fmt.Fprintf(&sb, "%d-", *spine.right)
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}

func (s *spine) num() int {
	var sb strings.Builder
	if s.left != nil {
		fmt.Fprint(&sb, *s.left)
	}
	fmt.Fprint(&sb, s.mid)
	if s.right != nil {
		fmt.Fprint(&sb, *s.right)
	}
	res, _ := strconv.Atoi(sb.String())
	return res
}

func parse(in []byte) *parsed {
	p := &parsed{}
	for i := 0; i < len(in); i++ {
		p.swords = append(p.swords, newSword(in, &i))
	}
	return p
}

func (p *parsed) checksum() int {
	res := 0
	for i := 0; i < len(p.swords); i++ {
		res += (i + 1) * p.swords[i].id
	}
	return res
}

func (p *parsed) String() string {
	var sb strings.Builder
	for _, sword := range p.swords {
		fmt.Fprintf(&sb, "%s\n", sword)
	}
	return sb.String()
}
