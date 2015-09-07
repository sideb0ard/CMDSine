package main

import "time"

func ticker(tickChan chan int) {
	// tickLength := (60000 / bpm) / 4 // 1min divided by bpm divided by 4 microticks
	tickLength := (60000 / bpm) // 1min divided by bpm divided by 4 microticks
	tickCounter := 1

	for {
		if bpm != 0 {
			//tickLength = (60000 / bpm) / 4
			tickLength = (60000 / bpm)
		}
		timer := time.NewTimer(time.Duration(tickLength) * time.Millisecond)
		tickChan <- tickCounter
		tickCounter++
		<-timer.C // pause before next iteration - length of bpm
	}
}
