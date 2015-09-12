package main

import (
	"fmt"
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

func newMixer() *mixer {
	m := &mixer{}
	return m
}

func (m *mixer) mix(signalChan chan *oscillator) {
	var err error
	m.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, m.processAudio)
	chk(err)
	defer m.Close()
	chk(m.Start())
	defer m.Stop()
	for s := range signalChan {
		if len(m.signals) > 5 {
			//m.signals[0].SilentStop()
			m.signals = m.signals[3:]
		}
		m.signals = append(m.signals, s)
	}
}

func (m *mixer) listChans() {
	for i, s := range m.signals {
		fmt.Println(i, s)
	}
}

func (m *mixer) processAudio(out [][]float32) {

	loopLength := 60000 / bpm / 60 * 16 // loop is 16 beats
	curPosition := math.Mod(float64(tickCounter), loopLength)

	for i := range out[0] {
		outval := float32(0)
		for _, s := range m.signals {
			ns := s.genNextSine()
			if s.amplitude.attack > 0 {
				attackFinish := loopLength * s.amplitude.attack
				decayStart := loopLength * 0.8
				if curPosition < attackFinish {
					adjustment := float32(curPosition / attackFinish)
					outval += adjustment * ns
				} else if curPosition > decayStart {
					adjustment := float32((loopLength - curPosition) / (loopLength - decayStart))
					outval += adjustment * ns
				} else {
					outval += ns
				}
			} else {
				outval += ns
			}
		}
		outval = outval / float32(len(m.signals))
		if outval < -1 || outval > 1 {
			fmt.Println("Oot", outval)
		}
		out[0][i] = outval
		out[1][i] = outval
	}
}
