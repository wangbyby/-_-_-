package heap

import (
	"go_code/alg/data/deferr"
)

//以下为 具体实现

//min heap
type Heap struct {
	Arr  []interface{} // 0...n-1
	size int           // 1...n
}

func NewHeap(r int) *Heap {
	return &Heap{Arr: make([]interface{}, r)}
}

func (h *Heap) BuildHeap() {
	h.size = len(h.Arr)
	for j := len(h.Arr) / 2; j >= 0; j-- {
		h.keepmin(j)
	}
}

func ReHeapSort(h *Heap) {
	h.BuildHeap()
	h.size = len(h.Arr)
	for i := len(h.Arr) - 1; i >= 1; i-- {
		h.Arr[0], h.Arr[i] = h.Arr[i], h.Arr[0]
		h.size--
		h.keepmin(0)
	}

}
func (h *Heap) less(a, b int) bool {
	return h.Arr[a] < h.Arr[b]
}

func (h *Heap) moveup(i int) { // n-1 --> 0
	p := (i - 1) / 2 // i 的父点
	if 0 <= p && h.less(i, p) {
		h.Arr[p], h.Arr[i] = h.Arr[i], h.Arr[p] // exch
		h.moveup(p)
	}
}

func (h *Heap) keepmin(i int) {
	l := i*2 + 1 // 左子
	r := i*2 + 2 //右子
	min := i

	if l < h.size && h.less(l, min) {
		min = l
	}
	if r < h.size && h.less(r, min) {
		min = r
	}

	if min != i {
		h.Arr[i], h.Arr[min] = h.Arr[min], h.Arr[i]
		h.keepmin(min)
	}
}
func (h *Heap) Push(v interface{}) {
	h.Arr = append(h.Arr, v)
	h.moveup(len(h.Arr) - 1)
}

func (h *Heap) Pop() (interface{}, error) {
	if len(h.Arr) <= 0 {
		return nil, deferr.ERR_BELOW_RANGE
	}
	res := h.Arr[0]
	h.Arr[0], h.Arr[h.size-1] = h.Arr[h.size-1], h.Arr[0]
	h.Arr = h.Arr[:h.size-1]
	h.size--
	h.keepmin(0)

	return res, nil
}
