package main

import (
	"fmt"
	hh "go_code/alg/data/heap"
	"math/rand"
)

type PlayHeap []int

func (h *PlayHeap) Swap(a, b int) {
	(*h)[a], (*h)[b] = (*h)[b], (*h)[a] // 或者 (*h)[a]
}

func (h *PlayHeap) Less(a, b int) bool { return (*h)[a] < (*h)[b] }
func (h *PlayHeap) Len() int {
	return len(*h)
}
func (h *PlayHeap) Getindex(i int) interface{} {
	return (*h)[i]
}
func (h *PlayHeap) Resetslice(a, b int) {
	*h = (*h)[a:b]
}
func (h *PlayHeap) Append(v interface{}) {
	*h = append(*h, v.(int))
}

func main() {
	var h PlayHeap = PlayHeap(make([]int, 10))
	for i := 0; i < 10; i++ {
		h[i] = rand.Intn(100)
	}

	hh.BuildHeap(&h)
	hh.HeapSort(&h)
	fmt.Println(h)
}
