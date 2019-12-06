package utils

import (
	"errors"
)

type Node struct {
	Weight              int
	Code                string
	Right, Left, Parent *Node
}
type Heap struct {
	Maxsize     int
	CurrentSize int
	List        []*Node
}

func (h *Heap) Init(size int) {
	h.List = make([]*Node, size)
	h.Maxsize = size
}
func (h *Heap) Insert(node *Node) (err error) {
	if h.Maxsize == h.CurrentSize {
		return errors.New("Heap Full")
	}
	h.List[h.CurrentSize] = node
	h.siftup(h.CurrentSize)
	h.CurrentSize++
	return
}
func (h *Heap) leftChild(n int) int {
	return 2*n + 1
}
func (h *Heap) siftup(n int) {
	i := n
	j := h.leftChild(n)
	tmp := h.List[j]
	for j < h.CurrentSize {
		if j < h.CurrentSize-1 && h.List[j].Weight > h.List[j+1].Weight {
			j++
		}
		if tmp.Weight > h.List[j].Weight {
			h.List[i] = h.List[j]
			i = j
			j = h.leftChild(j)
		} else {
			break
		}
	}
	h.List[i] = tmp

}
func (h *Heap) siftDown(n int) {
	indextmp := n
	tmpNode := h.List[n]
	for indextmp > 0 && h.List[(indextmp-1)/2].Weight > tmpNode.Weight {
		h.List[indextmp] = h.List[(indextmp-1)/2]
		indextmp = (indextmp - 1) / 2
	}
	h.List[indextmp] = tmpNode
}
func (h *Heap) RemoveMin() (min *Node, err error) {
	if h.CurrentSize == 0 {
		return nil, errors.New("Empty")
	} else {
		h.List[0], h.List[h.CurrentSize-1] = h.List[h.CurrentSize-1], h.List[0]
		h.CurrentSize--
		if h.CurrentSize > 1 {
			h.siftDown(0)
		}
		return h.List[h.CurrentSize], nil
	}
}
func (h *Heap) MergeTree(par, child1, child2 *Node) {
	par.Left = child1
	par.Right = child2
	par.Weight = child1.Weight + child2.Weight
}

type HuffmanTree struct {
	Root *Node
}

func (hm *HuffmanTree) BuildTree(weight []int, n int) {
	heap := &Heap{}
	heap.Init(n)
	var parent, firstch, secch *Node
	var NodeList []Node = make([]Node, n)
	for i := 0; i < n; i++ {
		NodeList[i].Weight = weight[i]
		heap.Insert(&NodeList[i])
	}
	for i := 0; i < n-1; i++ {
		parent = &Node{}

		firstch, _ = heap.RemoveMin()
		secch, _ = heap.RemoveMin()
		heap.MergeTree(parent, firstch, secch)
		heap.Insert(parent)
		hm.Root = parent
	}
}
func (hm *HuffmanTree) HuffmanCode() {
	hm.huffmanCode(hm.Root, "")
}

func (hm *HuffmanTree) huffmanCode(root *Node, code string) {
	root.Code = code
	hm.huffmanCode(root.Left, code+"0")
	hm.huffmanCode(root.Right, code+"1")
}
