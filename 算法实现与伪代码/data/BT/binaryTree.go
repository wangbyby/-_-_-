package BT

import "errors"

type Node struct {
	left, right, pre *Node
	Value            interface{}
}

type BinTree interface {
	Less(a, b interface{}) bool
}

//节点 , 以及用节点代替Tree
func NewTree() Node {
	return Node{}
}

func (node *Node) Insert(b BinTree, v interface{}) (err error) {
	return insertValue(b, node, v)
}
func (n *Node) Delete(b BinTree, v interface{}) (err error) {
	if n == nil {
		return errors.New("nil root")
	}
	ok, tmp := Search(b, n, v)
	if !ok {
		return errors.New("not contain this value")
	}
	p := tmp.pre
	if tmp.right == nil {
		p.left = tmp.left
		return
	} else {
		tt := tmp.right
		for tt.left != nil {
			tt = tt.left
		}
		tt.left = tmp.left
		return
	}
}
func Search(b BinTree, n *Node, v interface{}) (bool, *Node) {
	if n.Value == v {
		return true, n
	} else if b.Less(n.Value, v) {
		return Search(b, n.right, v)
	} else {
		return Search(b, n.left, v)
	}
}

//插入
func insertValue(b BinTree, n *Node, v interface{}) (err error) {
	if n.Value == nil && n != nil {
		n.Value = v
		return
	}
	switch b.Less(n.Value, v) {
	case true:
		if n.left == nil {
			tmp := &Node{Value: v}
			n.left = tmp
			tmp.pre = n
			return
		}
		insertValue(b, n.left, v)
	case false:
		if n.right == nil {
			tmp := &Node{Value: v}
			n.right = tmp
			tmp.pre = n
			return
		}
		insertValue(b, n.right, v)
	default:
		return errors.New("Insert binary tree err")
	}
	return
}
