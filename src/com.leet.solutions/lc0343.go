package com_leet_solutions

import (
	"fmt"
)

func Lc0343() {
	printTableHead(343)
	n := initLc343()
	printMidLine()
	res := integerBreak(n)
	fmt.Println(res)
	printTableTail()
}

func initLc343() int {
	n := 10
	fmt.Println(n)
	return n
}

func integerBreak(n int) int {
	res := make([]int, n + 1)
	for i:=0; i < 3 ; i++ {
		res[i] = 1
	}
	for i := 3; i <= n ; i++ {
		for j := 1; j < i; j++ {
			res[i] = max(res[i], max(j * (i - j), j * res[i - j]))
		}
	}
	return res[n]
}

func max (a int, b int) int {
	if a > b {
		return a
	}
	return b
}