package main

import (
	"fmt"
	"os"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func scalr(x int32) float32 {
	low := float32(-2147483647)
	high := float32(2147483647)
	r1 := float32(high - low)
	lscal := float32(-1)
	hscal := float32(1)
	r2 := hscal - lscal
	return (r2 / r1) * (float32(x) + (-1))
}

func help() {
	fmt.Println("HALP! my commands are: sine, exit, and jobbie")
}

func myexit() {
	fmt.Println("Later, dude...")
	os.Exit(0)
}
