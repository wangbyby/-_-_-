package main

import (
	"fmt"
	hea "go_code/alg/data/heap"
	"math/rand"
	"time"
)

const Limit = 20

func main() {

	//var cost time.Duration
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	count := 10
	//s := &s.SortHome{Limit: 30, A: make([]interface{}, count)}
	h := hea.NewHeap(count)
	// v := make([]int, count)
	// c := make([]int, count)
	for i := 0; i < count; i++ {
		//s.A[i] = r.Intn(count)
		//c[i], v[i] = r.Intn(count), r.Intn(count)
		h.Arr[i] = r.Intn(count)
	}
	h.BuildHeap()
	// for i := 0; i < count; i++ {
	// 	fmt.Printf("%v\t", h.Arr[i])
	// }

	fmt.Println(h.Arr)
	fmt.Println()
	for i := 0; i < 6; i++ {
		e, _ := h.Pop()
		fmt.Printf("%v\t", e)
	}
	for i := 0; i < 6; i++ {
		v := rand.Intn(count)
		fmt.Println(v)
		h.Push(v)
	}
	//hea.ReHeapSort(h)
	fmt.Println(h.Arr)
	// fmt.Println("混合")
	// start := time.Now()
	// QuickSort1(v, 0, len(v)-1)
	// insertSort(v, 0, len(v)-1)

	// cost = time.Since(start)
	// fmt.Printf("cost=[%s]\n", cost)
	// start := time.Now()
	// sort.Ints(c)
	// cost := time.Since(start)
	// fmt.Printf("cost=[%s]\n", cost)

	// fmt.Println("非并发")
	// cost = s.QuickSort()
	// fmt.Printf("cost=[%s]\n", cost)

	// fmt.Println("并发")
	// cost := s.QuickSort2()
	// fmt.Printf("cost=[%s]\n", cost)

	// for _, v := range s.A {
	// 	fmt.Printf("%v\t", v)
	// }
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

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func shellsort(items []int) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}

// func shellSort(buf []int) {
// 	//times := 0
// 	tmp := 0
// 	length := len(buf)
// 	incre := length
// 	for {
// 		incre /= 2
// 		for k := 0; k < incre; k++ {
// 			for i := k + incre; i < length; i += incre {
// 				for j := i; j > k; j -= incre {
// 					//times++
// 					if buf[j] < buf[j-incre] {
// 						tmp = buf[j-incre]
// 						buf[j-incre] = buf[j]
// 						buf[j] = tmp
// 					} else {
// 						break
// 					}
// 				}
// 			}
// 		}

// 		if incre == 1 {
// 			break
// 		}
// 	}
// 	//fmt.Println("shell's sort times: ", times)
// }

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
