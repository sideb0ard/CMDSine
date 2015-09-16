package main

import "math"

func newFM(signalChan chan SoundGen, carFreq, modFreq float64) {
	fm := &freqMod{carrier: &oscillator{vol: 0.6, freq: carFreq, phase: 0, amplitude: &envelope{attack: 0.2}},
		modulator: &oscillator{vol: 0.6, freq: modFreq, phase: 0, amplitude: &envelope{attack: 02}}, modAmp: freqRad * 100}
	signalChan <- fm
}

func (fm *freqMod) genNextSound() float32 {
	modValue := fm.modAmp * math.Sin(fm.modulator.phaseIncr)
	fm.carrier.phase += fm.carrier.phaseIncr + modValue
	if fm.carrier.phase >= twoPi {
		fm.carrier.phase -= twoPi
	}
	fm.modulator.phase += fm.modulator.phaseIncr
	if fm.modulator.phase >= twoPi {
		fm.modulator.phase -= twoPi
	}
	return float32(fm.carrier.vol * math.Sin(fm.carrier.phase))
}
