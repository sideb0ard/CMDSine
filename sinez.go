package main

import (
	"fmt"
	"math"
	"time"
)

func newSine(signalChan chan *oscillator, freq float64) {
	s := &oscillator{vol: 0.6, freq: freq, phase: 0, amplitude: &envelope{attack: 0}}
	signalChan <- s
}

func (o *oscillator) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz // Phase: %.2f // // AMP.attack: %.2f", o.vol, int(o.freq), o.phase, o.amplitude.attack)
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

func (o *oscillator) set(property string, val float64) {
	fmt.Println("Setting", property, "to", val)
	switch property {
	case "vol":
		fmt.Println("CHAnging VOL to", val)
		o.vol = float64(val)
	case "freq":
		fmt.Println("CHAnging FREQ to", val)
		o.freq = val
	case "attack":
		fmt.Println("CHAnging ATTACK to", val)
		o.amplitude.attack = val
	default:
		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
	}
}

func (o *oscillator) genNextSine() float32 {
	vol := float64(0)
	phaseIncr := o.freq * freqRad
	o.phase += phaseIncr
	if o.phase >= twoPi {
		o.phase -= twoPi
	}

	//fmt.Println("O.PHASE", o.phase)
	if o.amplitude.attack > 0 {
		//fmt.Println(vol)
		if o.phase < (twoPi * o.amplitude.attack) {
			// current location as percentage of the space between 0 and amp attack * vol
			//vol = o.vol * o.phase / (twoPi * o.amplitude.attack)
			vol = 0 * o.phase / (twoPi * o.amplitude.attack)
		} else {
			vol = o.vol
		}
		// fmt.Println(vol)
	} else {
		vol = o.vol
	}

	return float32(vol * math.Sin(o.phase))
}
