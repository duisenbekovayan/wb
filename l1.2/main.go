package main

import (
	"fmt"
	"sync"
)

type result struct {
	idx int
	n   int
	sq  int
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	resultsCh := make(chan result, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func(i, x int) {
			defer wg.Done()
			sq := x * x
			resultsCh <- result{
				idx: i,
				n:   x,
				sq:  sq,
			}
		}(i, n)
	}
	wg.Wait()
	out := make([]result, len(nums))
	for i := 0; i < len(nums); i++ {
		r := <-resultsCh
		out[r.idx] = r
	}
	for _, r := range out {
		fmt.Printf("%d^2 = %d\n", r.n, r.sq)
	}
}
