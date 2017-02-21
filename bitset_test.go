package bitset

import (
	"reflect"
	"testing"
)

var (
	arrs    = [][]int{[]int{}, []int{0}, []int{1}, []int{0, 1, 2}, []int{1, 4, 10}}
	bitsets = []Bitset{0, 1, 2, 7, 1042}
)

func TestCreate(t *testing.T) {
	for i, arr := range arrs {
		bs := NewFromSlice(arr)
		if bs != bitsets[i] {
			t.Fatalf("Incorrect bitset for %v, expected %d, got %d", arr, bitsets[i], bs)
		}
	}
}

func TestToSlice(t *testing.T) {
	for i, bs := range bitsets {
		arr := bs.ToSlice()
		if !reflect.DeepEqual(arr, arrs[i]) {
			t.Fatalf("Incorrect slice for bitset %d, expected %v, got %v", bs, arrs[i], arr)
		}
	}
}

func TestRemove(t *testing.T) {
	bs := NewFromSlice([]int{1, 4, 10})
	result := bs.RemoveMember(10)
	if !reflect.DeepEqual(result.ToSlice(), []int{1, 4}) {
		t.Fatalf("Incorrect result removing index. Got %d", result)
	}
}

func TestContains(t *testing.T) {
	bs := NewFromSlice([]int{0, 2, 4, 5, 7, 24})
	testMembers := []int{0, 1, 2, 15, 24, 25}
	answers := []bool{true, false, true, false, true, false}
	for i, v := range testMembers {
		if result := bs.Contains(v); result != answers[i] {
			t.Fatalf("Incorrect result with member %d, expected %v, got %v", v, answers[i], result)
		}
	}
}

func TestApplyMask(t *testing.T) {
	arr := []int{1, 2, 3}
	masks := []int{0, 1, 2, 3}
	answers := [][]int{[]int{}, []int{1}, []int{2}, []int{1, 2}}
	for i, mask := range masks {
		if result := applyMask(arr, mask); !reflect.DeepEqual(result, answers[i]) {
			t.Fatalf("Incorrect result from apply mask. Applying %d, expected %v, got %v", mask, answers[i], result)
		}
	}
}

func TestPowerSet(t *testing.T) {
	arr := []int{1, 2, 3}
	bs := NewFromSlice(arr)
	ps := bs.PowerSet()
	if len(ps) != len(arr)+1 {
		t.Fatalf("Incorrect length of powerset")
	}
	if ps[0][0] != NewFromSlice([]int{}) {
		t.Fatalf("Empty set not present in powersets of length 0")
	}
	if ps[len(arr)][0] != NewFromSlice([]int{1, 2, 3}) {
		t.Fatalf("Full set not present in powersets of length n")
	}
}

func TestUnion(t *testing.T) {
	bitsets := []Bitset{
		NewFromSlice([]int{}),
		NewFromSlice([]int{1}),
		NewFromSlice([]int{1, 2, 3}),
		NewFromSlice([]int{3}),
	}
	if u, a := bitsets[0].Union(bitsets[1]), NewFromSlice([]int{}); u != a {
		t.Fatalf("Incorrect union, expected %v, got %v", a, u)
	}
	if u, a := bitsets[1].Union(bitsets[2]), NewFromSlice([]int{1}); u != a {
		t.Fatalf("Incorrect union, expected %v, got %v", a, u)
	}
	if u, a := bitsets[1].Union(bitsets[3]), NewFromSlice([]int{}); u != a {
		t.Fatalf("Incorrect union, expected %v, got %v", a, u)
	}
	if u, a := bitsets[2].Union(bitsets[3]), NewFromSlice([]int{3}); u != a {
		t.Fatalf("Incorrect union, expected %v, got %v", a, u)
	}
}
