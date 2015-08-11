package main

import "code.google.com/p/portaudio-go/portaudio"

const sampleRate = 44100

type stereoSine struct {
	*portaudio.Stream
	time   float64 // counter
	vol    float32
	freqL  float64
	phaseL float64
	freqR  float64
	phaseR float64
}
