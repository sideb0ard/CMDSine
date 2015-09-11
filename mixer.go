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
	for i := range out[0] {
		outval := float32(0)
		curPosition := math.Mod(float64(tickCounter), loopLength)
		for _, s := range m.signals {
			ns := s.genNextSine()
			if s.amplitude.attack > 0 {
				attackFinish := loopLength * s.amplitude.attack
				decayStart := loopLength * 0.8
				if curPosition < attackFinish || curPosition > decayStart {
					//outval += float32(curPosition*attackFinish) * ns
					//fmt.Println("cur:", curPosition, "attackFinish:", attackFinish, "loopLength", loopLength)
					//fmt.Println("VOL SHOULD BE", curPosition/attackFinish)
					adjustment := float32(curPosition / attackFinish)
					if adjustment == 0 {
						fmt.Println("OFFT, MULTIPLY BY ZERO!")
					}
					outval += adjustment * ns
					//fmt.Println("TICK", tickCounter%int(loopLength), "LOOPLENGTH", loopLength)
				} else {
					outval += ns
				}
			} else {
				outval += ns
			}
		}
		out[0][i] = outval
		out[1][i] = outval
	}
}
