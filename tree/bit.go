package tree

import "golang.org/x/exp/constraints"

type BIT[integer constraints.Integer] struct {
	n   int
	bit []integer
}

func NewBIT[integer constraints.Integer](n int) *BIT[integer] {
	return &BIT[integer]{
		n:   n,
		bit: make([]integer, n+1),
	}
}

func (b *BIT[integer]) Update(idx int, diff integer) {
	for idx <= b.n {
		b.bit[idx] += diff
		idx += lowbit(idx)
	}
}

func (b *BIT[integer]) Query(idx int) integer {
	var res integer
	for idx > 0 {
		res += b.bit[idx]
		idx -= lowbit(idx)
	}
	return res
}

func (b *BIT[integer]) RangeQuery(l, r int) integer {
	return b.Query(r) - b.Query(l-1)
}

func lowbit(x int) int {
	return x & -x
}
