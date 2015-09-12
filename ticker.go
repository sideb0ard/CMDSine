package main

import "time"

func ticker(tickChan chan int) {
	//func ticker() {
	// tickLength := (60000 / bpm) / 4 // 1min divided by bpm divided by 4 microticks
	tickLength := (60000 / bpm) / 60 // 1min divided by bpm divided by 60 microticks
	tic := 1

	for {
		if bpm != 0 {
			tickLength = (60000 / bpm) / 60 // 60 Microticks per second
		}
		timer := time.NewTimer(time.Duration(tickLength) * time.Millisecond)

		// this is used for Prime Sub
		if tickCounter%60 == 0 {
			tic++
			select {
			case tickChan <- tic:
			default:
			}
		}

		tickCounter++
		<-timer.C // pause before next iteration - length of bpm
	}
}
