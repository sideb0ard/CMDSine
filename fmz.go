package main

import (
	"fmt"
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

func newFM(fmChan chan *FM, carFreq, modFreq float64) {
	forevs := make(chan bool)
	fm := newFMz(carFreq, modFreq, sampleRate)
	defer fm.Close()
	chk(fm.Start())
	defer fm.Stop()
	fmChan <- fm
	<-forevs
}

func newFMz(carFreq, modFreq, sampleRate float64) *FM {
	fm := &FM{car: &signalGenerator{amp: 0.6, freq: carFreq, phase: 0, phaseIncr: freqRad * carFreq},
		mod: &signalGenerator{amp: 0.6, freq: modFreq, phase: 0, phaseIncr: freqRad * modFreq}}
	var err error
	fm.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, fm.processAudio)
	chk(err)
	return fm
}

func (fm *FM) String() string {
	return fmt.Sprintf("FM:: CarFreq : %.2f // ModFreq: %.2f", fm.car.freq, fm.mod.freq)
}

func (fm *FM) processAudio(out [][]float32) {
	fm.mod.amp = freqRad * 100
	for i := range out[0] {
		// out[0][i] = float32(math.Sin(fm.carPhase * (fm.time / bpm)))
		// out[1][i] = float32(math.Sin(fm.carPhase * (fm.time / bpm)))
		out[0][i] = float32(math.Sin(fm.car.phase * float64(sumNum*int(bpm))))
		out[1][i] = float32(math.Sin(fm.car.phase * float64(sumNum*int(bpm))))
		fm.mod.freq = fm.mod.amp * math.Sin(fm.mod.phase)
		fm.car.phase = fm.car.phaseIncr + fm.mod.freq
		fm.mod.phase += fm.mod.phaseIncr
		fm.time++
	}
}
