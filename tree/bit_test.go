package tree

import "testing"

func TestBIT_INT(t *testing.T) {
	bit := NewBIT[int](5)
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)
	if bit.Query(1) != 1 {
		t.Error("error")
	}
	if bit.Query(2) != 3 {
		t.Error("error")
	}
	if bit.Query(3) != 6 {
		t.Error("error")
	}
	if bit.Query(4) != 10 {
		t.Error("error")
	}
	if bit.Query(5) != 15 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 3) != 6 {
		t.Error("error")
	}
	if bit.RangeQuery(2, 4) != 9 {
		t.Error("error")
	}
	if bit.RangeQuery(3, 5) != 12 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 5) != 15 {
		t.Error("error")
	}
}

func TestBIT_INT8(t *testing.T) {
	bit := NewBIT[int8](5)
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)
	if bit.Query(1) != 1 {
		t.Error("error")
	}
	if bit.Query(2) != 3 {
		t.Error("error")
	}
	if bit.Query(3) != 6 {
		t.Error("error")
	}
	if bit.Query(4) != 10 {
		t.Error("error")
	}
	if bit.Query(5) != 15 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 3) != 6 {
		t.Error("error")
	}
	if bit.RangeQuery(2, 4) != 9 {
		t.Error("error")
	}
	if bit.RangeQuery(3, 5) != 12 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 5) != 15 {
		t.Error("error")
	}
}

func TestBIT_UINT(t *testing.T) {
	bit := NewBIT[uint](5)
	bit.Update(1, 1)
	bit.Update(2, 2)
	bit.Update(3, 3)
	bit.Update(4, 4)
	bit.Update(5, 5)
	if bit.Query(1) != 1 {
		t.Error("error")
	}
	if bit.Query(2) != 3 {
		t.Error("error")
	}
	if bit.Query(3) != 6 {
		t.Error("error")
	}
	if bit.Query(4) != 10 {
		t.Error("error")
	}
	if bit.Query(5) != 15 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 3) != 6 {
		t.Error("error")
	}
	if bit.RangeQuery(2, 4) != 9 {
		t.Error("error")
	}
	if bit.RangeQuery(3, 5) != 12 {
		t.Error("error")
	}
	if bit.RangeQuery(1, 5) != 15 {
		t.Error("error")
	}
}
