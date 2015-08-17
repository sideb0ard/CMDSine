package main

import (
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

const (
	sampleRate = 44100
	twoPi      = math.Pi * 2
	freqRad    = twoPi / sampleRate
)

var bpm float64 = 94
var sumNum int = 0

type signalGenerator struct {
	amp       float64
	freq      float64
	phase     float64
	phaseIncr float64
}

type stereoSine struct {
	*portaudio.Stream
	time float64 // counter
	sine *signalGenerator
}

type FM struct {
	*portaudio.Stream
	time float64 // counter
	car  *signalGenerator
	mod  *signalGenerator
}
