package graph

import (
	"../deferr"
	"sort"
)

func (g *Graph) Dijkstra(s, t int) (p map[int]int, err error) {
	if s == t {
		return nil, deferr.Err_SOURCE_SAME_TARGET
	}
	d := make(map[int]float64, g.NumV)
	p = make(map[int]int, g.NumV)
	for i, _ := range d {
		d[i] = deferr.MaxWeight
	}
	d[s] = 0

	Q := SortBy(make([]Edge, g.NumV))

	for i := 0; i < g.NumV; i++ {
		e := Edge{From: -1, To: i, Weight: d[i]}
		Q[i] = e
	}

	for len(Q) != 0 {
		sort.Sort(Q)
		u := Q[0]
		Q = Q[1:]
		if err != nil {
			return nil, err
		}

		uu := u.To
		for v, _ := range g.AdjSet[uu] {
			ww, _ := g.W(uu, v)
			if d[v] > d[uu]+ww {
				d[v] = d[uu] + ww
				p[v] = uu
			}
		}
		for i, v := range d {
			Q[i].Weight = v
		}
	}
	return p, nil

}
