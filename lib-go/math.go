package ec

import (
	"fmt"
	"math/big"
)

func Abs[T ECSigned](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Mod[T ECSigned](n, m T) T {
	a := n % m
	if a < 0 {
		a += m
	}
	return a
}

func NeverToBigMod[T ECSigned](n, m T) T {
	if n < 0 {
		return n + m
	} else if n >= m {
		return n - m
	}
	return n
}

func Min[T ECInt](a ...T) T {
	min := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

func Max[T ECInt](a ...T) T {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}

func Abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func GCD(a, b int64) int64 {
	a = Abs(a)
	b = Abs(b)
	if a > b {
		a, b = b, a
	}
	for a != 0 {
		a, b = (b % a), a
	}
	return b
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func ModExp(b, e, m uint) uint {
	var res uint = 1
	for e > 0 {
		if (e % 2) == 1 {
			res = (res * b) % m
		}
		e = e / 2
		b = (b * b) % m
	}
	return res
}

func Product[T ECInt](ints ...T) T {
	var p T = 1
	for _, x := range ints {
		p *= x
	}
	return p
}

func Sum[T ECInt](ints ...T) T {
	var s T
	for _, x := range ints {
		s += x
	}
	return s
}

func EGCD[T ECInt](a, b T, x, y *T) T {
	if a == 0 {
		*x, *y = 0, 1
		return b
	}
	g := EGCD(b%a, a, x, y)
	*x, *y = *y-(b/a)*(*x), *x
	return g
}

func ChineseRemainderTheorem[T ECInt](an [][2]T) *T {
	p := an[0][1]
	for _, e := range an[1:] {
		p *= e[1]
	}
	var x T
	var y T
	var z T
	var j T
	for i := range an {
		a, n := an[i][0], an[i][1]
		q := p / n
		z = EGCD(n, q, &j, &y)
		if z != 1 {
			return nil
		}
		x += a * y * q
		for x < 0 {
			x += p
		}
		x %= p
	}
	return &x
}

var one = big.NewInt(1)

func CRT(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
