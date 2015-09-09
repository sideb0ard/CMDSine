package main

import (
	"fmt"

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

	for i := range out[0] {
		outval := float32(0)
		for _, s := range m.signals {
			outval += s.genNextSine()
		}
		// fmt.Println(outval)
		out[0][i] = outval
		out[1][i] = outval
	}
}
