package main

import (
	"fmt"
	"math"
	"time"
)

func newSine(signalChan chan *oscillator, freq float64) {
	s := &oscillator{vol: 0.6, freq: freq, phase: 0, phaseIncr: freq * freqRad}
	signalChan <- s
}

func (o *oscillator) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz // Phase: %.2f // PhIncr: %.2f // AMP.attack: %.2f", o.vol, int(o.freq), o.phase, o.phaseIncr, o.amplitude.attack)
}

func (o *oscillator) SilentStop() {
	curVol := o.vol
	for {
		if curVol > 0.2 {
			timer := time.NewTimer(time.Duration(300) * time.Millisecond)
			o.vol -= 0.2
			<-timer.C
		}
		return
	}
}

//
//func (g *oscillator) set(property string, val float64) {
//	fmt.Println("Setting ", property, " to ", val)
//	switch property {
//	case "vol":
//		fmt.Println("CHAnging VOL to", val)
//		g.sine.vol = float64(val)
//	case "freq":
//		fmt.Println("CHAnging FREQ to", val)
//		g.sine.freq = val
//	default:
//		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
//	}
//}

func (o *oscillator) genNextSine() float32 {
	vol := float64(0)
	o.phase += o.phaseIncr
	if o.phase >= twoPi {
		o.phase -= twoPi
	}
	if o.amplitude.attack > 0 {
		fmt.Println(vol)
		if o.phase < (twoPi * o.amplitude.attack) {
			vol = o.vol * o.phase * (twoPi * o.amplitude.attack)
		} else {
			vol = o.vol
		}
	} else {
		vol = o.vol
	}

	return float32(o.vol * math.Sin(o.phase))
}
