package com_leet_solutions

import (
"fmt"
)

func Lc0253() {
	printTableHead(253)
	intervals := initLc253()
	printMidLine()
	minRooms := minMeetingRooms(intervals)
	fmt.Println(minRooms)
	printTableTail()
}

func initLc253() [][]int {
	intervals := [][]int {{0,1},{2,4},{7,10}}
	fmt.Println(intervals)
	return intervals
}

func minMeetingRooms(intervals [][]int) int {
	timeline := generateTimeline(intervals)
	for _, v := range intervals {
		fillTimeline(v[0], v[1], timeline)
	}
	return findMaxOverlap(timeline)
}

func generateTimeline(intervals [][]int) []int {
	eof := 0
	for _, v := range intervals {
		if v[1] > eof {
			eof = v[1]
		}
	}
	return make([]int, eof + 1)
}

func fillTimeline(start int, end int, timeline []int){
	for i := start ; i < end; i ++ {
		timeline[i] += 1
	}
}

func findMaxOverlap(timeline []int) int {
	maxOL := 0
	for _,v := range timeline {
		if v > maxOL {
			maxOL = v
		}
	}
	return maxOL
}