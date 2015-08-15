package main

import "time"

func ticker(ticker chan int) {
	tickLength := (60000 / bpm) / 4 // 1min divided by bpm divided by 4 microticks

	for i := 0; ; i++ {
		if bpm != 0 {
			tickLength = (60000 / bpm) / 4
		}
		timer := time.NewTimer(time.Duration(tickLength) * time.Millisecond)
		ticker <- i
		<-timer.C // pause before next iteration - length of bpm
	}
}
