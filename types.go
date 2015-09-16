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

var bpm float64 = 47
var tickCounter = 1
var tic = 1 // used for Prime Sub
var sumNum int = 0
var tickLength = 60000 / bpm / 60
var loopLength = tickLength * 16 // loop is 16 beats

type oscillator struct {
	vol       float64
	freq      float64
	phase     float64
	phaseIncr float64
	clock     float64

	amplitude *envelope
}

type freqMod struct {
	carrier   *oscillator
	modulator *oscillator
	modAmp    float64
}

type envelope struct {
	attack    float64
	sustain   float64
	decay     float64
	increment float64
}

type mixer struct {
	*portaudio.Stream
	signals []SoundGen
}

type SoundGen interface {
	genNextSound() float32
}
