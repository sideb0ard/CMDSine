package main

import "time"

func ticker(tickChan chan int) {

	for {
		timer := time.NewTimer(time.Duration(tickLength) * time.Millisecond)

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
