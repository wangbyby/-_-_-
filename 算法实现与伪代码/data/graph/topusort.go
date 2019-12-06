package graph

import (
	"../sq"
)

/*
topusort(G)
	queue s
	reult r
	Q = G.V
	for i in Q
		if i.in == 0
			s.add(i)
	while s != nil
		e = s.pop()
		r.add(e)
		for  i  in adj(e)
			i.count--
			if i.count ==0
				s.push(i)
	return r


*/

func (g *Graph) TopuSort() (res []int, err error) {
	res = make([]int, g.NumV)
	tmp := make(map[int]int)
	for i, v := range g.de {
		tmp[i] = v.in
	}
	q := sq.NewQueue()

	for i, _ := range g.AdjSet {

		if g.de[i].in == 0 {
			q.Push(i)
		}
	}

	for q.Len > 0 {
		e := q.Pop().(int)
		res[e] = e
		for i, v := range g.Adj(e) {
			tmp[v.To]--
			if tmp[v.To] == 0 {
				q.Push(i)
			}
		}
		return res, nil
	}
}
