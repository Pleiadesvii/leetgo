package com_leet_solutions

import (
	"fmt"
)

func Lc0046() {
	printTableHead(46)
	nums := initLc46()
	printMidLine()
	numsList := permute(nums)
	fmt.Println(numsList)
	printTableTail()
}

func initLc46() []int {
	nums := []int {1,2,3}
	fmt.Println(nums)
	return nums
}

func permute(nums []int) [][]int {
	res := [][]int {}
	path := make([]int, len(nums))
	visited := make([]bool, len(nums))
	res = dfs(res, 0, nums, path, visited)
	return res
}

func dfs(res [][]int, index int, nums []int, path []int, visited []bool) [][]int {
	if index >= len(nums) {
		tmp := []int {}
		for _,v := range path {
			tmp = append(tmp, nums[v])
		}
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			path[index] = i
			visited[i] = true
			res = dfs(res, index + 1, nums, path, visited)
			visited[i] = false
		}
	}
	return res
}