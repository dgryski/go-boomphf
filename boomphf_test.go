package boomphf

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dgryski/go-radixsort"
	"github.com/dustin/go-humanize"
)

func testKeys(t *testing.T, n int) {

	rand.Seed(0)

	t0 := time.Now()
	t.Logf("generating keys")
	k := make([]uint64, 0, n)
	for i := 1; i <= n; i++ {
		k = append(k, uint64(i))
	}
	t.Logf("generated %d keys in %v", n, time.Since(t0))

	t0 = time.Now()
	h := New(2, k)

	t.Logf("construct(%v)=%v", n, time.Since(t0))
	sz := h.Size()
	t.Logf("size=%v (%.2f bits per item)", humanize.Bytes(uint64(sz)), float64(8*sz)/float64(n))

	got := make([]uint64, 0, n)

	t0 = time.Now()
	for _, v := range k {
		r := h.Query(v)
		got = append(got, r)
	}
	took := time.Since(t0)
	t.Logf("query(%v)=%v, %v per item", n, took, took/time.Duration(n))

	radixsort.Uint64s(got)

	for i, v := range got {
		if v != uint64(i+1) {
			t.Fatalf("failed: v=%v i+1=%v", v, i+1)
		}
	}

}

func TestKeys(t *testing.T) { testKeys(t, 1e7) }
