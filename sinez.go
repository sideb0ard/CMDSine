package main

import (
	"fmt"
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

//var sinez []*stereoSine

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
	s := &stereoSine{sine: &signalGenerator{amp: 0.6, freq: freq, phase: 0, phaseIncr: freq * freqRad}}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *stereoSine) String() string {
	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz", g.sine.amp, int(g.sine.freq))
}

func (g *stereoSine) set(property string, val float64) {
	fmt.Println("Setting ", property, " to ", val)
	switch property {
	case "vol":
		fmt.Println("CHAnging VOL to", val)
		g.sine.amp = float64(val)
	case "freq":
		fmt.Println("CHAnging FREQ to", val)
		g.sine.freq = val
	default:
		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
	}
}

func (g *stereoSine) processAudio(out [][]float32) {
	// melodyting
	// scale := 1.58

	// SAWTOOTH
	// sawIncr := 2 * g.freq / sampleRate
	// sawVal := float64(-1)

	// triable
	//twoDivPi := 2.0 / math.Pi

	// square
	// midPoint := twoPi *

	for i := range out[0] {
		// value := float64(0)
		// for p := float64(1); p <= 4; p++ {
		// 	value += float64(math.Sin(g.phase*p)) / p
		// }
		// g.phase += g.phaseIncr
		// if g.phase >= twoPi {
		// 	g.phase -= twoPi
		// }
		// out[0][i] = g.vol * float32(value/scale*g.time/bpm)
		// out[1][i] = g.vol * float32(value/scale*g.time/bpm)

		//out[0][i] = g.vol * float32(math.Sin(g.phase*(g.time/bpm)))
		//out[1][i] = g.vol * float32(math.Sin(g.phase*(g.time/bpm)))
		out[0][i] = float32(g.sine.amp * math.Sin(g.sine.phase*bpm))
		out[1][i] = float32(g.sine.amp * math.Sin(g.sine.phase*bpm))
		g.sine.phase = g.sine.phase + g.sine.phaseIncr

		// simple sine
		//out[0][i] = g.vol * float32(math.Sin(g.phase))
		//out[1][i] = g.vol * float32(math.Sin(g.phase))
		//g.phase = g.phase + g.phaseIncr
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

		// triangle wave
		// triVal := g.phase * twoDivPi
		// if triVal < 0 {
		// 	triVal = 1.0 + triVal
		// } else {
		// 	triVal = 1.0 - triVal
		// }
		// out[0][i] = g.vol * float32(triVal)
		// out[1][i] = g.vol * float32(triVal)

		// square

		// g.phase += g.phaseIncr
		// if g.phase >= twoPi {
		// 	g.phase -= twoPi
		// }

		g.time++
	}
}
