package com_leet_solutions

import (
	"fmt"
)

func Lc0047() {
	printTableHead(47)
	nums := initLc47()
	printMidLine()
	numsList := permuteUnique(nums)
	fmt.Println(numsList)
	printTableTail()
}

func initLc47() []int {
	nums := []int {1,1,3,3}
	fmt.Println(nums)
	return nums
}

func permuteUnique(nums []int) [][]int {
	res := [][]int {}
	path := make([]int, len(nums))
	visited := make([]bool, len(nums))
	res = dfs2(res, 0, nums, path, visited)
	return res
}

func dfs2(res [][]int, index int, nums []int, path []int, visited []bool) [][]int {
	if index >= len(nums) {
		tmp := []int {}
		for _,v := range path {
			tmp = append(tmp, nums[v])
		}
		res = checkSameAndInsert(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			path[index] = i
			visited[i] = true
			res = dfs2(res, index + 1, nums, path, visited)
			visited[i] = false
		}
	}
	return res
}

func checkSameAndInsert(res [][]int, tmp []int) [][]int {
	hasSame := false
	for _,v := range res {
		i := 0
		for i = 0; i < len(v); i++ {
			if tmp[i] != v[i] {
				break
			}
		}
		if i == len(v) {
			hasSame = true
		}
	}
	if !hasSame {
		res = append(res, tmp)
	}
	return res
}