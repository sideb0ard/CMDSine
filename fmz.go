package main

import (
	"fmt"
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

func newFM(fmChan chan *FM, carFreq, modFreq float64) {
	forevs := make(chan bool)
	fm := newFMz(carFreq, modFreq)
	defer fm.Close()
	chk(fm.Start())
	defer fm.Stop()
	fmChan <- fm
	<-forevs
}

func newFMz(carFreq, modFreq float64) *FM {
	fm := &FM{car: &oscillator{vol: 0.6, freq: carFreq, phase: 0, phaseIncr: freqRad * carFreq},
		mod: &oscillator{vol: 0.6, freq: modFreq, phase: 0, phaseIncr: freqRad * modFreq}}
	var err error
	fm.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, fm.processAudio)
	chk(err)
	return fm
}

func (fm *FM) String() string {
	return fmt.Sprintf("FM:: CarFreq : %.2f // ModFreq: %.2f", fm.car.freq, fm.mod.freq)
}

func (fm *FM) processAudio(out [][]float32) {
	modAmp := freqRad * float64(100)
	modVal := modAmp * math.Sin(fm.mod.phase)
	for i := range out[0] {
		out[0][i] = float32(math.Sin(fm.car.phase))
		out[1][i] = float32(math.Sin(fm.car.phase))
		modVal = modAmp * math.Sin(fm.mod.phase)
		fm.car.phase += fm.car.phaseIncr + modVal
		fm.mod.phase += fm.mod.phaseIncr
	}
}
