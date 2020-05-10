package com_leet_solutions

import (
	"fmt"
)

func Lc1197() {
	printTableHead(1197)
	x, y := initLc1197()
	printMidLine()
	minStep := minKnightMoves(x, y)
	fmt.Println(minStep)
	printTableTail()
}

func initLc1197() (int, int){
	target := []int {142,68}
	fmt.Println(target)
	return target[0], target[1]
}

func minKnightMoves(x int, y int) int {
	xt := absInt(x) + 2
	yt := absInt(y) + 2
	startP := []int {2,2}
	dirA := [][]int {{1,2},{2,1},{2,-1},{1,-2},{-1,-2},{-2,-1},{-2,1},{-1,2}}

	rightBD := maxInt(xt, startP[0]) + 2
	topBD := maxInt(yt, startP[1]) + 2
	positionA := [][]int {}
	positionA = append(positionA, startP)
	visitedA := generateVisitedMap(rightBD,topBD)
	visitedA[startP[0]][startP[1]] = true

	step := 0
	for len(positionA) > 0 {
		for size := len(positionA); size > 0; size-- {
			curP := positionA[0]
			positionA = positionA[1:]
			if curP[0] == xt && curP[1] == yt {
				return step
			}
			for _, dir := range dirA {
				nextP := []int {0,0}
				nextP[0] = curP[0] + dir[0]
				nextP[1] = curP[1] + dir[1]
				if (nextP[0] <= rightBD && nextP[0] >= 0 && nextP[1] <= topBD && nextP[1] >= 0 && !visitedA[nextP[0]][nextP[1]]) {
					visitedA[nextP[0]][nextP[1]] = true
					positionA = append(positionA, nextP)
				}
			}
		}
		step++
	}

	return -1
}


func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateVisitedMap(x int, y int) [][]bool {
	var visited [][]bool
	for i := 0; i < x + 1 ;  i++ {
		var visitedX []bool
		for j := 0; j < y + 1 ; j++ {
			visitedX = append(visitedX, false)
		}
		visited = append(visited, visitedX)
	}
	return visited
}