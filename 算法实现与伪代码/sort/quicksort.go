package sort

import (
	"sync"
	"time"
)

type SortHome struct {
	ss    sync.WaitGroup //全局锁
	Limit int            // 划分quicksort 与 insertsort的界限
	A     []interface{}  // 数据切片
}

func (s *SortHome) Len() int                   { return len(s.A) }
func (s *SortHome) Less(a, b interface{}) bool { return a.(int) < b.(int) }

//非并发
func (s *SortHome) QuickSort() time.Duration {
	start := time.Now()
	s.quicksort1(s.A, 0, s.Len()-1)
	cost := time.Since(start)
	return cost
}

//并发 并返回运行时间
func (s *SortHome) QuickSort2() time.Duration {
	start := time.Now()
	s.ss.Add(1)
	s.quicksort(s.A, 0, s.Len()-1)
	s.ss.Wait()
	cost := time.Since(start)
	return cost
}
func (s *SortHome) quicksort(A []interface{}, p, r int) {
	defer s.ss.Done()
	if p < r {
		if r-p < s.Limit {
			s.insertSort(A, p, r)
		} else {
			q := s.part(A, p, r)
			s.ss.Add(2)
			go s.quicksort(A, q+1, r)
			go s.quicksort(A, p, q-1)
		}
	}
}
func (s *SortHome) insertSort(A []interface{}, p, r int) {
	for j := p + 1; j <= r; j++ {
		key := A[j]
		i := j - 1
		for i >= p && s.Less(key, A[i]) {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}
func (s *SortHome) part(A []interface{}, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if s.Less(A[j], x) {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}
func (s *SortHome) quicksort1(A []interface{}, p, r int) {
	if p < r {
		if r-p < s.Limit {
			s.insertSort(A, p, r)
		} else {
			q := s.part(A, p, r)
			s.quicksort1(A, q+1, r)
			s.quicksort1(A, p, q-1)
		}
	}
}

func (s *SortHome) MoreQuickSort() time.Duration {
	start := time.Now()
	s.morequicksort(0, s.Len()-1)
	s.insertSort(s.A, 0, s.Len()-1)
	return time.Since(start)
}

func (s *SortHome) morequicksort(p, r int) {
	if p < r {
		if r-p < s.Limit {
			s.insertSort(s.A, p, r)
		} else {
			q := s.part(s.A, p, r)
			go s.morequicksort(p, q-1)
			go s.morequicksort(q+1, r)
		}
	}
}

// func part(A []int, p, r int) int {
// 	x := A[r]
// 	i := p - 1
// 	for j := p; j < r; j++ {
// 		if A[j] < x {
// 			i++
// 			A[i], A[j] = A[j], A[i]
// 		}
// 	}
// 	A[i+1], A[r] = A[r], A[i+1]
// 	return i + 1
// }

// func insertSort(A []int, p, r int) {
// 	for j := p + 1; j <= r; j++ {
// 		key := A[j]
// 		i := j - 1
// 		for i >= p && A[i] > key {
// 			A[i+1] = A[i]
// 			i--
// 		}
// 		A[i+1] = key
// 	}

// }

// func QuickSort1(A []int, p, r int) {

// 	if p < r {
// 		if r-p < 30 { //30是试出来的
// 			insertSort(A, p, r)
// 		} else {
// 			q := part(A, p, r)

// 			go QuickSort1(A, q+1, r)
// 			go QuickSort1(A, p, q-1)
// 		}
// 	}
// }
