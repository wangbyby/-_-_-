package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := name{Task: Doing}
	n.Task()
}

func Doing() {
	fmt.Println("i am doing sth")
}

type name struct {
	Task func()
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
