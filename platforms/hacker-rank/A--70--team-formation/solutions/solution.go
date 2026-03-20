package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func solve(skills []int) int {
	if len(skills) == 0 {
		return 0
	}

	sort.Ints(skills)

	end := map[int][]int{}

	push := func(heap []int, v int) []int {
		heap = append(heap, v)
		i := len(heap) - 1
		for i > 0 && heap[(i-1)/2] > heap[i] {
			heap[(i-1)/2], heap[i] = heap[i], heap[(i-1)/2]
			i = (i - 1) / 2
		}
		return heap
	}

	pop := func(heap []int) ([]int, int) {
		v := heap[0]
		last := heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		if len(heap) > 0 {
			heap[0] = last
			i := 0
			for {
				l := 2*i + 1
				r := 2*i + 2
				s := i
				if l < len(heap) && heap[l] < heap[s] {
					s = l
				}
				if r < len(heap) && heap[r] < heap[s] {
					s = r
				}
				if s == i {
					break
				}
				heap[i], heap[s] = heap[s], heap[i]
				i = s
			}
		}
		return heap, v
	}

	for _, x := range skills {
		if h, ok := end[x-1]; ok && len(h) > 0 {
			var length int
			end[x-1], length = pop(h)

			end[x] = push(end[x], length+1)
		} else {
			end[x] = push(end[x], 1)
		}
	}

	ans := int(^uint(0) >> 1)

	for _, heap := range end {
		for _, v := range heap {
			if v < ans {
				ans = v
			}
		}
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		skills := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &skills[i])
		}

		fmt.Println(solve(skills))
	}
}
