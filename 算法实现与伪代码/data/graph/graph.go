package graph

import (
	"errors"
	"fmt"

	"../ownerr"
	"../set"
)

/*
	字段:

		点集 //用数字表示点 0...n
		邻接表 (顺便储存边集)
		距离
	函数:
		权重w
		距离d
*/
type Edge struct {
	From, To int     // 表示边
	Weight   float64 //权重
}

type SortBy []Edge

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].Weight < a[j].Weight }

type degrees struct {
	in, out int
}
type Graph struct {
	//入度 出度
	de map[int]degrees
	//邻接表
	AdjSet map[int]map[int]Edge

	//每个点的距离 用map存储  插入/删除
	//Distance map[int]float64

	NumV int // 顶点数目
	NumE int //边数

}

func (g *Graph) Show() {
	for i, v := range g.AdjSet {
		for ii, vv := range v {
			fmt.Printf("起始点:%v,终点:%v,权重:%v\n", i, ii, vv.Weight)
		}
	}
}

func (g *Graph) Adj(u int) map[int]Edge {
	return g.AdjSet[u]
}

func NewGraph() (g Graph) {
	g = Graph{}
	g.AdjSet = make(map[int]map[int]Edge, 1)
	g.de = make(map[int]degrees, 0)
	g.NumV, g.NumE = 0, 0
	return g
}

// i 为起点, j为终点 边的权重
func (g *Graph) W(i, j int) (float64, error) {
	kk, ok := g.AdjSet[i]
	if !ok {
		return 0, errors.New("empty source")
	}
	kkk, ok := kk[j]
	if !ok {
		return 0, errors.New("empty end")
	}
	return kkk.Weight, nil
}

func (g *Graph) InsertEdge(e Edge) {
	mapk, ok := g.AdjSet[e.From]
	if !ok {
		// 没有该起点 就加上该点
		g.AdjSet[e.From] = make(map[int]Edge, 1)
		g.AdjSet[e.From][e.To] = e
		g.NumV++
		g.NumE++
		return
	}
	//有该点 就插入
	mapk[e.To] = e
	g.NumE++
}

func (g *Graph) DeleteEdge(e Edge) (err error) {
	//只删边
	kk, ok := g.AdjSet[e.From]
	if !ok {
		return ownerr.ERR_EDGE_FROM
	}
	_, ok = kk[e.To]
	if !ok {
		return ownerr.ERR_EDGE_TO
	}
	delete(kk, e.To)
	g.NumE--
	return
}

func (g *Graph) CopyGraph() (cg *Graph) {
	cg = NewGraph()
	cg.NumE, cg.NumV = g.NumE, g.NumV
	// cg.Distance = make(map[int]float64)
	// for i, v := range g.Distance {
	// 	cg.Distance[i] = v
	// }

	cg.VerSet = set.NewSetHash()
	for i, v := range g.VerSet.Set {
		cg.VerSet.Set[i] = v
	}

	cg.AdjSet = make(map[int]map[int]Edge)
	for i, v := range g.AdjSet {
		cg.AdjSet[i] = make(map[int]Edge)
		for ii, vv := range v {
			cg.AdjSet[i][ii] = vv
		}
	}
	return cg
}
