package main

import (
	"fmt"

	"code.google.com/p/portaudio-go/portaudio"
)

func newMixer() *mixer {
	m := &mixer{}
	return m
}

//func mixingDesk(signalChan chan *oscillator) {
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

// func newMixingDesk(sinezChan chan *stereoSine) {

func (m *mixer) processAudio(out [][]float32) {

	// scale := float64(1 / 214748)
	for i := range out[0] {
		//fmt.Println(i)

		// sine
		outval := float32(0)
		//scale /= float32(len(m.signals))
		for _, s := range m.signals {
			// fmt.Println("Scale", scale)
			outval += s.genNextSine()
		}
		out[0][i] = outval
		out[1][i] = outval

	}
}
