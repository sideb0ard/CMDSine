package main

import "code.google.com/p/portaudio-go/portaudio"

const sampleRate = 44100

var bpm float64 = 120

type stereoSine struct {
	*portaudio.Stream
	time   float64 // counter
	vol    float32
	freqL  float64
	phaseL float64
	freqR  float64
	phaseR float64
}

type FM struct {
	car stereoSine // carrier
	mod stereoSine // modulator
}
