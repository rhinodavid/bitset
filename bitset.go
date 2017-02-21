package bitset

import "log"

// Bitset represents a set of positive integers as a single integer. For each integer
// in the set, the corresponding bit in the resultant Bitset to 1.
type Bitset int

func pow(base, exp int) (r int) {
	if exp == 0 {
		return 1
	}
	r = base
	for i := 1; i < exp; i++ {
		r *= base
	}
	return
}

// NewFromSlice takes a slice of positive integers and returns the corresponding
// Bitset
func NewFromSlice(a []int) (r Bitset) {
	for i, v := range a {
		if v < 0 {
			log.Fatalf("All values in the slice must be positive. Saw %d at index %d", v, i)
		}
		r = r + Bitset(pow(2, v))
	}
	return
}

// ToSlice returns an array of integers representing the members of the set
// contained in the Bitset
func (b Bitset) ToSlice() []int {
	x := b
	r := make([]int, 0)
	for i := 0; x > 0; i++ {
		if x&1 == 1 {
			r = append(r, i)
		}
		x = x >> 1
	}
	return r
}

// RemoveMember takes an integer value
func (b Bitset) RemoveMember(i int) (r Bitset) {
	if i < 0 {
		log.Fatalf("Attempt to remove negative member from the Bitset: %d", i)
	}
	// check to see if index exists
	if b&Bitset(pow(2, i)) == Bitset(pow(2, i)) {
		r = b
		return r - Bitset(pow(2, i))
	}
	return b
}

// Contains accepts a positive integer and returns true if it is
// contained in the Bitset and false if it is not
func (b Bitset) Contains(i int) bool {
	if i < 0 {
		log.Fatalf("Attempt to check negative member in Bitset: %d", i)
	}
	return b&Bitset(pow(2, i)) == Bitset(pow(2, i))
}

func countBits(x int) (r int) {
	for y := x; y > 0; y = y >> 1 {
		if y&1 == 1 {
			r++
		}
	}
	return
}

func applyMask(a []int, m int) []int {
	r := make([]int, 0)
	n := len(a)
	for q := 0; m > 0; q++ {
		if m&1 == 1 {
			if q < n {
				r = append(r, a[q])
			}
		}
		m = m >> 1
	}
	return r
}

// PowerSet takes operates on a Bitset and returns an array
// where the index is the number of items in the powersets
// and the values are slices of Bitsets representing those
// powersets
func (b Bitset) PowerSet() [][]Bitset {
	set := b.ToSlice()
	n := len(set)
	result := make([][]Bitset, n+1)
	mask := pow(2, n) - 1
	for ; mask >= 0; mask-- {
		subSet := applyMask(set, mask)
		bs := NewFromSlice(subSet)
		ssLen := len(subSet)
		if result[ssLen] == nil {
			result[ssLen] = []Bitset{bs}
		} else {
			result[ssLen] = append(result[ssLen], bs)
		}
	}
	return result
}

// Union operates on one Bitset, accepts a second Bitset,
// and returns the union of the two
func (b Bitset) Union(c Bitset) Bitset {
	return b & c
}
