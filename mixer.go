package main

import "code.google.com/p/portaudio-go/portaudio"

func mixingDesk(signalChan chan *oscillator) {
	m := &mixer{}
	var err error
	m.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, m.processAudio)
	chk(err)
	defer m.Close()
	chk(m.Start())
	defer m.Stop()
	for s := range signalChan {
		if len(m.signals) > 5 {
			m.signals = m.signals[2:]
		}
		m.signals = append(m.signals, s)
	}
}

// func newMixingDesk(sinezChan chan *stereoSine) {

func (m *mixer) processAudio(out [][]float32) {

	for i := range out[0] {

		// sine
		outval := float32(0)
		for _, s := range m.signals {
			outval += s.genNextSine()
		}
		out[0][i] = outval
		out[1][i] = outval

	}
}
