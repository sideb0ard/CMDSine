package main

import (
	"fmt"
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

//var sinez []*stereoSine

func newSine(sinezChan chan *stereoSine, freq float64) {
	forever := make(chan bool)
	s := newStereoSine(freq, freq, sampleRate)
	defer s.Close()
	chk(s.Start())
	defer s.Stop()
	sinezChan <- s
	<-forever
}

func newStereoSine(freqL, freqR, sampleRate float64) *stereoSine {
	s := &stereoSine{nil, 0, 0.6, freqL / sampleRate, 0, freqR / sampleRate, 0}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *stereoSine) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz", g.vol, int(g.freqL*sampleRate))
}

func (g *stereoSine) set(property string, val float64) {
	fmt.Println("Setting ", property, " to ", val)
	switch property {
	case "vol":
		fmt.Println("CHAnging VOL to", val)
		g.vol = float32(val)
	case "freq":
		fmt.Println("CHAnging FREQ to", val)
		g.freqL = val / sampleRate
		g.freqR = val / sampleRate
	default:
		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
	}
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		// fmt.Println("FREQ ", g.stepL, g.stepR)
		out[0][i] = g.vol * float32(math.Sin(2*math.Pi*g.phaseL*(g.time/190)))
		_, g.phaseL = math.Modf(g.phaseL + g.freqL)
		out[1][i] = g.vol * float32(math.Sin(2*math.Pi*g.phaseR*(g.time/190)))
		_, g.phaseR = math.Modf(g.phaseR + g.freqR)
		g.time++
	}
}
