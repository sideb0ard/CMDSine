package main

import "math"

func newSine(signalChan chan *oscillator, freq float64) {
	s := &oscillator{vol: 0.6, freq: freq, phase: 0, phaseIncr: freq * freqRad}
	signalChan <- s
}

//func (g *stereoSine) String() string {
//	return fmt.Sprintf("SINE:: Vol: %.2f // Freq: %d Hz", g.sine.vol, int(g.sine.freq))
//}
//
//func (g *stereoSine) SilentStop() {
//	curVol := g.sine.vol
//	for {
//		if curVol > 0.2 {
//			timer := time.NewTimer(time.Duration(300) * time.Millisecond)
//			g.sine.vol -= 0.2
//			<-timer.C
//		}
//
//		g.Close()
//		return
//	}
//}
//
//func (g *stereoSine) set(property string, val float64) {
//	fmt.Println("Setting ", property, " to ", val)
//	switch property {
//	case "vol":
//		fmt.Println("CHAnging VOL to", val)
//		g.sine.vol = float64(val)
//	case "freq":
//		fmt.Println("CHAnging FREQ to", val)
//		g.sine.freq = val
//	default:
//		fmt.Println("SMOKE WEED EVERYDAYAYAYY")
//	}
//}
//
//func (g *stereoSine) processAudio(out [][]float32) {
//
//	for i := range out[0] {
//
//		// sine
//		out[0][i] = float32(g.sine.vol * math.Sin(g.sine.phase))
//		out[1][i] = float32(g.sine.vol * math.Sin(g.sine.phase))
//		g.sine.phase += g.sine.phaseIncr
//		if g.sine.phase >= twoPi {
//			g.sine.phase -= twoPi
//		}
//
//		// sawtooth
//		//out[0][i] = g.vol * float32(sawVal)
//		//out[1][i] = g.vol * float32(sawVal)
//		//sawVal += sawIncr
//		//if sawVal >= 1 {
//		//	sawVal -= 2
//		//}
//	}
//}

func (o *oscillator) genNextSine() float32 {
	o.phase += o.phaseIncr
	if o.phase >= twoPi {
		o.phase -= twoPi
	}
	return float32(o.vol * math.Sin(o.phase))
}
