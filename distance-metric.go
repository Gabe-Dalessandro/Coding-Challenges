package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1,2,1,1,2,3}
	ans := distanceMetric(arr) // should be [5,3,3,4,3,0]
	fmt.Println(ans)
}


func distanceMetric(arr []int) []int {
	ans := make([]int, len(arr))
	m := make(map[int][]int)

	// create the map
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		_, present := m[val]
		if present {
			m[val] = append(m[val], i)
		} else {
			m[val] = []int{i}
		}
	}

	for i := 0; i < len(arr); i++ {
		idxs := m[arr[i]]
		for j := 0; j < len(idxs); j++ {
			ans[i] += int(math.Abs(float64(idxs[j]-i)))
		}
	}

	return ans
}