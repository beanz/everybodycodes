package main

import (
	"fmt"
	"io"
	"math"
	"strings"

	. "github.com/beanz/everybodycodes/lib-go"
)

func parts(i1, i2, i3 []byte) (string, string, string) {
	p1 := parse(i1, false)
	p2 := parse(i2, false)
	p3 := parse(i3, true)
	return p1.String(), p2.String(), p3.String()
}

func main() {
	p1, p2, p3 := parts(Input(1, ""), Input(2, ""), Input(3, ""))
	fmt.Printf("Part 1: %v\nPart 2: %v\nPart 3: %v\n", p1, p2, p3)
}

type (
	parsed struct {
		left, right *node
	}
	node struct {
		left, right *node
		v, id       int
		ch          byte
	}
)

func (n *node) find(id int) []*node {
	if n == nil {
		return nil
	}
	var r []*node
	if n.id == id {
		r = append(r, n)
	}
	if f := n.left.find(id); f != nil {
		r = append(r, f...)
	}
	if f := n.right.find(id); f != nil {
		r = append(r, f...)
	}
	return r
}

func (n *node) add(o *node) {
	if o.v < n.v {
		if n.left == nil {
			n.left = o
		} else {
			n.left.add(o)
		}
	} else {
		if n.right == nil {
			n.right = o
		} else {
			n.right.add(o)
		}
	}
}

func (n *node) String() string {
	if n == nil {
		return ""
	}
	var sb strings.Builder
	n.str(&sb, "")
	return sb.String()
}

func (n *node) str(w io.Writer, prefix string) {
	if n == nil {
		return
	}
	fmt.Fprintf(w, "%s%c%d\n", prefix, n.ch, n.v)
	n.left.str(w, prefix+"  ")
	n.right.str(w, prefix+"  ")
}

func (n *node) depthString(d int, ds map[int]string) {
	if n == nil {
		return
	}
	ds[d] += fmt.Sprintf("%c", n.ch)
	n.left.depthString(d+1, ds)
	n.right.depthString(d+1, ds)
}

func parse(in []byte, part3 bool) *parsed {
	p := &parsed{}
	for i := 0; i < len(in); i++ {
		if in[i] == 'S' {
			var id int
			i, id = ChompUInt[int](in, i+5)
			fn := p.find(id)
			ln := fn[0]
			rn := fn[1]
			ln.ch, ln.v, rn.ch, rn.v = rn.ch, rn.v, ln.ch, ln.v
			if part3 {
				ln.left, ln.right, rn.left, rn.right =
					rn.left, rn.right, ln.left, ln.right
			}
			Dg("L", p.left)
			Dg("R", p.right)
			continue
		}
		var id, l, r int
		i, id = ChompUInt[int](in, i+7)
		i, l = ChompUInt[int](in, i+7)
		lch := in[i+1]
		i, r = ChompUInt[int](in, i+11)
		rch := in[i+1]
		i += 3
		Df("%d %d%c %d%c\n", id, l, lch, r, rch)
		ln := &node{v: l, id: id, ch: lch}
		if p.left == nil {
			p.left = ln
		} else {
			p.left.add(ln)
		}
		rn := &node{v: r, id: id, ch: rch}
		if p.right == nil {
			p.right = rn
		} else {
			p.right.add(rn)
		}
	}
	return p
}

func (p *parsed) find(id int) []*node {
	var r []*node
	r = append(r, p.left.find(id)...)
	r = append(r, p.right.find(id)...)
	return r
}

func (p *parsed) String() string {
	longest := func(ds map[int]string) string {
		m := ""
		id := math.MaxInt
		for k, v := range ds {
			if len(v) > len(m) {
				m = v
				id = k
			}
			if len(v) == len(m) && k < id {
				m = v
				id = k
			}
		}
		return m
	}
	ds := map[int]string{}
	p.left.depthString(0, ds)
	left := longest(ds)
	ds = map[int]string{}
	p.right.depthString(0, ds)
	right := longest(ds)
	return left + right
}
