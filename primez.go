package main

import (
	"fmt"
	"math/rand"
	"time"
)

func primez(sinezChan chan *stereoSine, sinez *[]*stereoSine) {

	tic := make(chan int)

	bassprimez := []int{53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	midprimez := []int{101, 103, 107, 109, 113,
		127, 131, 137, 139, 149, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
		233, 239, 241, 257, 263}

	highprimez := []int{269, 271, 277, 281,
		283, 293, 307, 311, 313, 317, 331, 337, 347, 349,
		353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457, 461, 463,
		467, 479, 487, 491}

	// ADDITIVE
	go gen(sinezChan, sinez, tic, 3, highprimez)
	go gen(sinezChan, sinez, tic, 11, bassprimez)
	go gen(sinezChan, sinez, tic, 7, midprimez)

	// REMOVITIVE
	go rem(sinez, 5)

	tickLength := 60000 / bpm
	tickCounter := 1
	for {
		if bpm != 0 {
			tickLength = 60000 / bpm
		}
		timer := time.NewTimer(time.Duration(tickLength) * time.Millisecond)
		tic <- tickCounter
		tickCounter++
		<-timer.C
	}

}

func gen(sinezChan chan *stereoSine, sinez *[]*stereoSine, tic chan int, primeTicker int, primeRange []int) {

	var nom int
	rand.Seed(42)
	for {
		nom = <-tic
		if nom%primeTicker == 0 {
			randyFreq := primeRange[rand.Intn(len(primeRange))]
			fmt.Printf("OOH[%d] -> got one: %d - choosing freq %d\n", primeTicker, nom, randyFreq)
			go newSine(sinezChan, float64(randyFreq))
			*sinez = append(*sinez, <-sinezChan)

		}
	}

}

func rem(sinez *[]*stereoSine, primeTicker int) {
	rand.Seed(477)
	for {
		fmt.Println("LEN(SINEZ)", len(*sinez))
		timer := time.NewTimer(time.Duration(500) * time.Millisecond)

		if len(*sinez) > primeTicker {
			sineNumToRemove := rand.Intn(len(*sinez))
			randyRemoval := (*sinez)[sineNumToRemove]
			fmt.Println("Oooh, gonna silence sine ", randyRemoval)
			randyRemoval.SilentStop()
			*sinez = append((*sinez)[sineNumToRemove:], (*sinez)[sineNumToRemove+1:]...)
		}
		<-timer.C
	}
}
