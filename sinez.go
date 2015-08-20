package main

import (
	"fmt"
	"math"
	"time"

	"code.google.com/p/portaudio-go/portaudio"
)

func newSine(sinezChan chan *stereoSine, freq float64) {
	forever := make(chan bool)
	s := newStereoSine(freq)
	defer s.Close()
	chk(s.Start())
	defer s.Stop()
	sinezChan <- s
	<-forever
}

func newStereoSine(freq float64) *stereoSine {
	s := &stereoSine{sine: &oscillator{vol: 0.6, freq: freq, phase: 0, phaseIncr: freq * freqRad}}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *stereoSine) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz", g.sine.vol, int(g.sine.freq))
}

func (g *stereoSine) SilentStop() {
	curVol := g.sine.vol
	for {
		if curVol > 0.2 {
			timer := time.NewTimer(time.Duration(300) * time.Millisecond)
			g.sine.vol -= 0.2
			<-timer.C
		}

		g.Close()
		return
	}
}

func (g *stereoSine) set(property string, val float64) {
	fmt.Println("Setting ", property, " to ", val)
	switch property {
	case "vol":
		fmt.Println("CHAnging VOL to", val)
		g.sine.vol = float64(val)
	case "freq":
		fmt.Println("CHAnging FREQ to", val)
		g.sine.freq = val
	default:
		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
	}
}

func (g *stereoSine) processAudio(out [][]float32) {

	for i := range out[0] {

		// sine
		out[0][i] = float32(g.sine.vol * math.Sin(g.sine.phase))
		out[1][i] = float32(g.sine.vol * math.Sin(g.sine.phase))
		g.sine.phase += g.sine.phaseIncr
		if g.sine.phase >= twoPi {
			g.sine.phase -= twoPi
		}

		// sawtooth
		//out[0][i] = g.vol * float32(sawVal)
		//out[1][i] = g.vol * float32(sawVal)
		//sawVal += sawIncr
		//if sawVal >= 1 {
		//	sawVal -= 2
		//}
	}
}
