package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Array []int

func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Array) Len() int {
	return len(a)
}
func (a Array) Less(i, j int) bool {
	return a[i] < a[j]
}

func generateSlice(size int) Array {

	slice := make(Array, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
func TestSort(t *testing.T) {
	a := generateSlice(10)
	Sort(a)
	if !sort.IsSorted(a) {
		t.Error("sort erro")
	}
}
