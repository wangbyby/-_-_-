package Graph

type Edge struct {
	From, To int
	weight   float64
}

type Graph struct {
	AdjGraph map[int][]Edge // 邻接表

	EdgeSet map[int]Edge // 边集

	V int //点的数量
	E int //边的数量
}
