package main

import (
	"fmt"

	"code.google.com/p/portaudio-go/portaudio"
)

func newMixer() *mixer {
	m := &mixer{}
	return m
}

func (m *mixer) mix(signalChan chan SoundGen) {
	var err error
	m.Stream, err = portaudio.OpenDefaultStream(0, 2, sampleRate, 0, m.processAudio)
	chk(err)
	defer m.Close()
	chk(m.Start())
	defer m.Stop()
	for s := range signalChan {
		if len(m.signals) > 5 {
			//go func() { fmt.Println("yar!") }() // s m.signals[0]) { s.SilentStop() }
			go func(m *mixer) {
				fmt.Println("SILENCIO!")
				// m.signals[0].SilentStop()
				m.signals = m.signals[1:]
			}(m) // s m.signals[0]) { s.SilentStop() }
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

	// curPosition := math.Mod(float64(tickCounter), loopLength)
	// // attackFinish := loopLength * s.amplitude.attack
	// attackFinish := loopLength * 0.2
	// decayStart := loopLength * 0.8
	// attack_adjustment := float32(curPosition / attackFinish)
	// decay_adjustment := float32((loopLength - curPosition) / (loopLength - decayStart))

	for i := range out[0] {

		outval := float32(0)
		for _, s := range m.signals {

			ns := s.genNextSound()

			// NEED A CASE HERE
			// if s.amplitude.attack > 0 {
			// 	if curPosition < attackFinish {
			// 		outval += attack_adjustment * ns
			// 	} else if curPosition > decayStart {
			// 		outval += decay_adjustment * ns
			// 	} else {
			// 		outval += ns
			// 	}
			// } else {
			// 	outval += ns
			// }
			outval += ns

		}
		outval = outval / float32(len(m.signals))
		if outval < -1 || outval > 1 {
			fmt.Println("Oot", outval)
		}

		out[0][i] = outval
		out[1][i] = outval
	}
}
