package main

import (
	"fmt"
	"math/rand"
)

func primez(signalChan chan *oscillator, tickChan chan int) {

	forever := make(chan bool)

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
	go gen(signalChan, tickChan, 3, highprimez)
	go gen(signalChan, tickChan, 11, bassprimez)
	go gen(signalChan, tickChan, 7, midprimez)
	go firstSig(signalChan, bassprimez)

	<-forever

}

func firstSig(signalChan chan *oscillator, primeRange []int) {
	rand.Seed(4)
	randyFreq := primeRange[rand.Intn(len(primeRange))]
	// fmt.Printf("Innnnniiiit!", randyFreq)
	newSine(signalChan, float64(randyFreq))
}

func gen(signalChan chan *oscillator, tic chan int, primeTicker int, primeRange []int) {

	var nom int
	rand.Seed(42)
	for {
		nom = <-tic
		if nom%primeTicker == 0 {
			randyFreq := primeRange[rand.Intn(len(primeRange))]
			fmt.Printf("OOH[%d] -> got one: %d - choosing freq %d\n", primeTicker, nom, randyFreq)
			newSine(signalChan, float64(randyFreq))
		}
	}

}

// func rem(sinez *[]*stereoSine, primeTicker int) {
// 	rand.Seed(477)
// 	for {
// 		fmt.Println("LEN(SINEZ)", len(*sinez))
// 		timer := time.NewTimer(time.Duration(500) * time.Millisecond)
//
// 		if len(*sinez) > primeTicker {
// 			sineNumToRemove := rand.Intn(len(*sinez))
// 			randyRemoval := (*sinez)[sineNumToRemove]
// 			fmt.Println("Oooh, gonna silence sine ", randyRemoval)
// 			randyRemoval.SilentStop()
// 			*sinez = append((*sinez)[sineNumToRemove:], (*sinez)[sineNumToRemove+1:]...)
// 		}
// 		<-timer.C
// 	}
// }
