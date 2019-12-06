package sq

// stack queue
import "errors"

type Node struct {
	next  *Node
	Value interface{}
}

type Queue struct {
	len        int
	head, tail *Node
}

type Stack struct {
	top *Node
	len int
}

func NewHeap() Queue {
	return Queue{len: 0, head: nil, tail: nil}
}

/*
	Stack示意图
		* --> * --> *
		^
		|
		top
*/
func NewStack() Stack {
	return Stack{len: 0, top: nil}
}

func (s *Stack) Push(v interface{}) {
	if s.len == 0 {
		s.len++
		s.top = &Node{Value: v}
		return
	}
	tmp := &Node{Value: v}
	tmp.next = s.top
	s.top = tmp
	s.len++
}

func (s *Stack) PopElement() (interface{}, error) {
	if s.len <= 0 {
		return nil, errors.New("Empty Stack")
	}
	s.len--
	tmp := s.top
	s.top = s.top.next
	return tmp.Value, nil
}

func (q *Queue) Push(v interface{}) error {
	if q.len == 0 {
		q.tail = &Node{Value: v}
		q.head = q.tail
		q.len++
		return nil
	} else {
		tmp := &Node{Value: v}
		q.tail = tmp
		q.len++
		return nil
	}
}

func (q *Queue) Pop() (v interface{}, err error) {
	if q.len < 1 {
		return nil, errors.New("empry queue")
	}
	q.len--
	tmp := q.head
	q.head = q.head.next
	return tmp.Value, nil
}

func (q *Queue) Len() int { return q.len }
