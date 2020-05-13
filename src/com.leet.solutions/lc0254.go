package com_leet_solutions

import (
	"fmt"
	"math"
)

func Lc0254() {
	printTableHead(254)
	n := initLc254()
	printMidLine()
	factors := getFactors(n)
	fmt.Println(factors)
	printTableTail()
}

func initLc254() int {
	n := 32
	fmt.Println(n)
	return n
}

func getFactors(n int) [][]int {
	factors := [][]int {}
	for i := 2; i <= int(math.Sqrt(float64(n))) ; i++ {
		if checkTimeResult(n, i, n/i) {
			f := []int {i, n/i}
			factors = checkMonoAndAppend(factors, f)
			for _,v := range getFactors(n/i) {
				f_in := []int {i}
				f_in = append(f_in, v...)
				factors = checkMonoAndAppend(factors, f_in)
			}

		}
	}

	return factors
}

func checkTimeResult(tar int, a int,b int) bool {
	if tar == a * b {
		return true
	}
	return false
}

func checkMonoAndAppend(nodes [][]int, node []int) [][]int{
	old := node[0]
	for _,v := range node {
		if v < old {
			return nodes
		}
		old = v
	}
	return append(nodes, node)
}
