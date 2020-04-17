package com_leet_solutions

import (
	"fmt"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

func Lc0002() {
	printTableHead(2)
	initLc()
	solution()
	printTableTail()
}

func solution() {
	fmt.Println("here is the solution of lc0002")
	log.Info("info")
}

func initLc() {
	fmt.Println("here is init of lc0002")
}
