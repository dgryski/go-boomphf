package boomphf

import (
	"math/rand"
	"sort"
	"testing"
)

func testKeys(t *testing.T, n int) {

	rand.Seed(0)

	var k []uint64

	for i := 1; i <= n; i++ {
		k = append(k, uint64(i))
	}

	h := New(3, k)

	t.Logf("h=%b %+v", h.b, h.ranks)

	var got []int

	for _, v := range k {
		r := h.Query(v)
		t.Logf("lookup(%v)=%v", v, r)
		got = append(got, int(r))
	}

	sort.Ints(got)

	for i, v := range got {
		if v != (i + 1) {
			t.Fatalf("failed: v=%v i+1=%v", v, i+1)
		}
	}

}

func TestKeys(t *testing.T) { testKeys(t, 1000) }
