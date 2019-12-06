package Graph

import (
	"fmt"
)

var (
	s       [][]int
	visited []bool
	p       []int
	Ind     []int
	count   int
)

func Init(n, m int) { //初始化
	var a, b int
	p = make([]int, 0, n)
	Ind = make([]int, n)

	for i := 0; i < n; i++ { //初始化访问数组
		visited = append(visited, false)
	}

	for i := 0; i < n; i++ {
		s1 := make([]int, 0, n)
		s = append(s, s1)
	}

	for i := 0; i < m; i++ {
		fmt.Scanln(&a, &b)
		s[a] = append(s[a], b)
		Ind[b]++
	}
}

func Topolopy_Sort(n int) bool { //拓扑排序
	count := 0
	var count1 int
	flag := true
	for {
		if count < n {
			count1 = count
			for i := 0; i < n; i++ {
				if visited[i] == false && Ind[i] == 0 {
					p = append(p, i)
					visited[i] = true
					count++
					fmt.Println(count)
					length := len(s[i])
					for j := 0; j < length; j++ {
						Ind[s[i][j]]--
					}
				}
			}
		} else {
			break
		}
		if count1 == count && count != n-1 {
			flag = false
			break
		}
	}

	return flag
}
