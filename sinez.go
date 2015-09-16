package main

import (
	"fmt"
	"math"
	"time"
)

func newSine(signalChan chan SoundGen, freq float64) {
	s := &oscillator{vol: 0.6, freq: freq, phase: 0, phaseIncr: freq * freqRad, amplitude: &envelope{attack: 0.2}}
	signalChan <- s
}

func (o *oscillator) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz // Phase: %.2f // // AMP.attack: %.2f", o.vol, int(o.freq), o.phase, o.amplitude.attack)
}

func (o *oscillator) SilentStop() {
	curVol := o.vol
	for {
		if curVol > 0 {
			timer := time.NewTimer(time.Duration(400) * time.Millisecond)
			o.vol -= 0.1
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

func (o *oscillator) genNextSound() float32 {
	o.phase += o.phaseIncr
	if o.phase >= twoPi {
		o.phase -= twoPi
	}
	return float32(o.vol * math.Sin(o.phase))
}
