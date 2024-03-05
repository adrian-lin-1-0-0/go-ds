package tree

import "golang.org/x/exp/constraints"

// BIT is a binary indexed tree.
type BIT[integer constraints.Integer] []integer

// NewBIT returns a new BIT with n elements.
func NewBIT[integer constraints.Integer](n int) BIT[integer] {
	return make([]integer, n+1)
}

// Update adds diff to the idx-th element.
func (b BIT[integer]) Update(idx int, diff integer) {
	for idx <= len(b)-1 {
		b[idx] += diff
		idx += lowbit(idx)
	}
}

// Query returns the sum of the first idx elements.
func (b BIT[integer]) Query(idx int) integer {
	var res integer
	for idx > 0 {
		res += b[idx]
		idx -= lowbit(idx)
	}
	return res
}

// RangeQuery returns the sum of the elements in the range [l, r].
func (b BIT[integer]) RangeQuery(l, r int) integer {
	return b.Query(r) - b.Query(l-1)
}

// lowbit returns the lowest bit of x.
func lowbit(x int) int {
	return x & -x
}
