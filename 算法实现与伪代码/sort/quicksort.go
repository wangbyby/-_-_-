package sort

func generateSlice(size int) Array {

	slice := make(Array, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//希尔排序...不懂(＠_＠;)
//shell sort
func Shellsort(items []int) {
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

type Array []int

func (a Array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Array) Len() int {
	return len(a)
}
func (a Array) Less(i, j int) bool {
	return a[i] < a[j]
}

//只对向量
type Interface interface {
	Swap(i, j int)
	Len() int
	Less(a, b int) bool
}

func Sort(qs Interface) {
	quicksort1(qs, 0, qs.Len()-1)
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
func part(qs Interface, p, r int) int {
	//若以r为主键下标
	i := p - 1
	for j := p; j < r; j++ {
		if qs.Less(j, r) {
			i++
			qs.Swap(i, j)
		}
	}
	qs.Swap(i+1, r)
	return i + 1
}

func quicksort1(A Interface, p, r int) {
	if p < r {
		if r-p < 20 {
			insertSort(A, p, r)
		} else {
			q := part(A, p, r)
			quicksort1(A, q+1, r)
			quicksort1(A, p, q-1)
		}
	}
}

func quickSort(qs Interface, p, r int) {
	if p < r {
		q := part(qs, p, r)
		quickSort(qs, q+1, r)
		quickSort(qs, p, q-1)
	}
}

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

func insertSort(qs Interface, p, r int) {
	for j := p + 1; j <= r; j++ {
		for i := j; i > p && qs.Less(i, i-1); i-- {
			qs.Swap(i, i-1)
		}

	}
}

func HeapSort(data Interface) {
	buildHeap(data)
	size := data.Len()
	for i := size - 1; i >= 1; i-- {
		data.Swap(0, i)
		size--
		siftUp(data, 0, size)
	}
	//for循环完是从大到小

	//转换为从小到大
	size = data.Len()
	for i := 0; i < size>>1; i++ {
		data.Swap(i, size-1-i)
	}
}

//最小堆 minheap
func BuildHeap(data Interface) {
	lenData := data.Len()
	for i := data.Len() >> 1; i >= 0; i-- {
		siftUp(data, i, lenData)
	}
}

func siftUp(data Interface, i, size int) {
	min := i
	l := i<<1 + 1
	r := (i + 1) << 1
	if l < size && data.Less(l, i) {
		min = l
	}
	if r < size && data.Less(r, min) {
		min = r
	}
	if min != i {
		data.Swap(min, i)
		siftUp(data, min, size)
	}
}
