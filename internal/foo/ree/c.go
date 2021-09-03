package ree

import "sort"

var arr = []int{1, 5, 6, 2, 7, 3, 7, 2}

func Run() {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]<arr[j]
	})
}
