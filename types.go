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

var bpm float64 = 120

type stereoSine struct {
	*portaudio.Stream
	time      float64 // counter
	vol       float32
	freq      float64
	phase     float64
	phaseIncr float64
}

type FM struct {
	*portaudio.Stream
	time     float64 // counter
	carFreq  float64
	carIncr  float64
	carPhase float64
	carAmp   float64
	modFreq  float64
	modIncr  float64
	modPhase float64
	modAmp   float64
}
