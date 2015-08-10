package main

import (
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

//var sinez []*stereoSine

func newSine(sinezChan chan *stereoSine) {
	forever := make(chan bool)
	s := newStereoSine(256, 256, sampleRate)
	defer s.Close()
	chk(s.Start())
	defer s.Stop()
	sinezChan <- s
	<-forever
}

func newStereoSine(freqL, freqR, sampleRate float64) *stereoSine {
	s := &stereoSine{nil, 0, freqL / sampleRate, 0, freqR / sampleRate, 0}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *stereoSine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phaseL * (g.time / 190)))
		_, g.phaseL = math.Modf(g.phaseL + g.stepL)
		out[1][i] = float32(math.Sin(2 * math.Pi * g.phaseR * (g.time / 190)))
		_, g.phaseR = math.Modf(g.phaseR + g.stepR)
		g.time++
	}
}
