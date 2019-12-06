package main

import (
	"fmt"
	"go_code/alg/sort"
	"math/rand"
	"time"
)

const Limit = 20

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 10000000
	s := &sort.SortHome{Limit: 30, A: make([]interface{}, count)}
	v := make([]int, count)
	for i := 0; i < s.Len(); i++ {
		s.A[i] = r.Intn(count)
		v[i] = r.Intn(count)
	}
	// fmt.Println("非并发")
	// cost := s.QuickSort()
	// fmt.Printf("cost=[%s]\n", cost)

	fmt.Println("并发")
	cost := s.QuickSort2()
	fmt.Printf("cost=[%s]\n", cost)

	// for _, v := range s.A {
	// 	fmt.Printf("%v\t", v)
	// }
	fmt.Println("混合")
	start := time.Now()
	QuickSort1(v, 0, len(v)-1)
	insertSort(v, 0, len(v)-1)
	cost = time.Since(start)
	fmt.Printf("cost=[%s]\n", cost)
	// for _, v := range v {
	// 	fmt.Printf("%v\t", v)
	// }
}

func part(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] < x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func insertSort(A []int, p, r int) {
	for j := p + 1; j <= r; j++ {
		key := A[j]
		i := j - 1
		for i >= p && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}

func QuickSort1(A []int, p, r int) {

	if p < r {
		if r-p < Limit {
			insertSort(A, p, r)
		} else {
			q := part(A, p, r)

			go QuickSort1(A, q+1, r)
			go QuickSort1(A, p, q-1)
		}
	}
}

// func quicksort1(A []int, p, r int) {
// 	if p < r {
// 		if r-p < Limit {
// 			insertSort(A, p, r)
// 		} else {
// 			q := part(A, p, r)
// 			quicksort1(A, q+1, r)
// 			quicksort1(A, p, q-1)
// 		}
// 	}
// }
