package Graph

import "math"

func (g *Graph) BellmanFord(s int) bool {

	d := make(map[int]float64, 0)
	d[s] = 0
	for i, _ := range g.EdgeSet {
		d[i] = math.Inf(0)
	}

	for i := 1; i < g.V-1; i++ {
		for _, v := range g.EdgeSet {
			if d[v.To] > d[v.From]+v.weight {
				d[v.To] = d[v.From] + v.weight
			}
		}
	}

	for _, v := range g.EdgeSet {
		if d[v.To] > d[v.From]+v.weight {
			return false
		}
	}
	return true
}
