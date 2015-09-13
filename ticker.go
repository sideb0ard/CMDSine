package main

import "time"

func ticker(tickChan chan int) {
	tickLength := (60000 / bpm) / 60 // 60 mictoticks per beat
	tic := 1

	for {
		if bpm != 0 {
			tickLength = (60000 / bpm) / 60
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
		<-timer.C
	}
}
