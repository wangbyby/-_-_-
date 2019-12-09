package heap

import (
	"errors"
)

type HeapInter interface {
	Less(a, b int) bool // 下标访问
	Swap(a, b int)      // 下标访问
	Len() int           // 数组长度
	Getindex(i int) interface{}
	Resetslice(a, b int) // 暂时没有好名字
	Append(v interface{})
}

func BuildHeap(h HeapInter) {
	size := h.Len()
	for j := h.Len() / 2; j >= 0; j-- {
		keepMin(h, j, size)
	}
}

func HeapSort(h HeapInter) {
	BuildHeap(h)
	size := h.Len()
	for i := h.Len() - 1; i >= 1; i-- {
		h.Swap(0, i)
		size--
		keepMin(h, 0, size)
	}

}

func moveUp(h HeapInter, i int) {
	p := (i - 1) / 2 // i 的父点
	if 0 <= p && h.Less(i, p) {
		h.Swap(p, i)
		moveUp(h, p)
	}
}
func keepMin(h HeapInter, i, size int) {
	l := i*2 + 1 // 左子
	r := i*2 + 2 //右子
	min := i

	if l < size && h.Less(l, min) {
		min = l
	}
	if r < size && h.Less(r, min) {
		min = r
	}

	if min != i {
		h.Swap(i, min)
		keepMin(h, min, size)
	}
}
func Pop(h HeapInter) (interface{}, error) {
	if h.Len() <= 0 {
		return nil, errors.New("Bad Len")
	}
	size := h.Len() - 1
	res := h.Getindex(0)
	h.Resetslice(0, size)
	h.Swap(0, size)
	keepMin(h, 0, size)
	return res, nil
}
func Push(h HeapInter, v interface{}) {
	h.Append(v)
	moveUp(h, h.Len()-1)
}
