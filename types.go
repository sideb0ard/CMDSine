package main

import "code.google.com/p/portaudio-go/portaudio"

const sampleRate = 44100

type stereoSine struct {
	*portaudio.Stream
	time          float64 // counter
	stepL, phaseL float64
	stepR, phaseR float64
}