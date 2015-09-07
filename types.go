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

type oscillator struct {
	vol       float64
	freq      float64
	phase     float64
	phaseIncr float64

	amplitude envelope
}

type envelope struct {
	attack    float64
	sustain   float64
	decay     float64
	increment float64
}

type mixer struct {
	*portaudio.Stream
	signals []*oscillator
}
