package Graph

import (
	"math"
	"sort"
)

func (g *Graph) Prim(r int) {
	//r 为初始节点
	for _, u := range g.EdgeSet {
		u.weight = math.Inf(1)
		u.From = -1
	}
	tmp := g.EdgeSet[r]
	tmp.weight = 0
	g.EdgeSet[r] = tmp
	var Q map[int]Edge
	for i, v := range g.EdgeSet {
		Q[i] = v
	}

	for len(Q) != 0 {
		u := ExMin(Q).From
		for _, v := range g.AdjGraph[u] {
			val, ok := Q[v.To]
			if ok && val.weight < v.weight {
				v.From = u
				v.weight = val.weight
			}
		}
	}
}

func ExMin(Q map[int]Edge) Edge {
	e := Edge{From: 0, To: 0, weight: math.Inf(1)}
	for _, v := range Q {
		if v.weight < e.weight {
			e = v
		}
	}
	delete(Q, e.From)
	return e
}

//点 用 int 表示
//边 struct from to key
//边集 map visited

var visited = make(map[int]bool, 0)

type SortBy []Edge

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return Compare(a[i], a[j]) }
func Compare(b, c Edge) bool        { return b.weight > c.weight }

func (g *Graph) K() map[int]Edge {
	//最小生成树
	t := make(map[int]Edge, 0)

	//复制一个边集
	var copyDedgeSet SortBy = make([]Edge, g.V)
	for i, v := range g.EdgeSet {
		copyDedgeSet[i] = v
	}
	sort.Sort(copyDedgeSet) //逆序 取最后为最小
	for i := g.V - 1; i > -1; i-- {
		if !visited[copyDedgeSet[i].To] {
			t[i] = copyDedgeSet[i]
			visited[copyDedgeSet[i].To] = true
		}
	}
	return t
}
