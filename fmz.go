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
	fm := &FM{nil, 0, carFreq, freqRad * carFreq, 0, 0.6, modFreq, freqRad * modFreq, 0, 0.6}
	var err error
	fm.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, fm.processAudio)
	chk(err)
	return fm
}

func (fm *FM) String() string {
	return fmt.Sprintf("FM:: CarFreq : %.2f // ModFreq: %.2f", fm.carFreq, fm.modFreq)
}

func (fm *FM) processAudio(out [][]float32) {
	fm.modAmp = freqRad * 100
	for i := range out[0] {
		// out[0][i] = float32(math.Sin(fm.carPhase * (fm.time / bpm)))
		// out[1][i] = float32(math.Sin(fm.carPhase * (fm.time / bpm)))
		out[0][i] = float32(math.Sin(fm.carPhase * float64(sumNum)))
		out[1][i] = float32(math.Sin(fm.carPhase * float64(sumNum)))
		fm.modFreq = fm.modAmp * math.Sin(fm.modPhase)
		fm.carPhase = fm.carIncr + fm.modFreq
		fm.modPhase += fm.modIncr
		fm.time++
	}
}
