package graph

import "sort"

func (g *Graph) Prim(s int) []Edge {

	visited := make(map[int]bool, g.NumV)
	Q := SortBy(make([]Edge, 0))
	for from, v := range g.AdjSet {
		for to, vv := range v {
			var e Edge
			e.From = from
			e.To = to
			e.Weight = vv.Weight
			Q = append(Q, e)
		}
	}
	mst := make([]Edge, 0)
	sort.Sort(Q)
	visited[s] = true
	for _, v := range Q {
		if !visited[v.From] || !visited[v.To] {
			visited[v.From], visited[v.To] = true, true
			mst = append(mst, v)
		}
	}
	return mst
}
