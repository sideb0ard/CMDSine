package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"

	"code.google.com/p/portaudio-go/portaudio"
)

func main() {
	//cmds := []string{"ls", "exit", "jobbie"}

	PS2 := ansi.Color("#CMDSine> ", "magenta")

	signalChan := make(chan *oscillator)

	tickChan := make(chan int)
	go ticker(tickChan)
	// go fib(tickerChan)

	portaudio.Initialize()
	defer portaudio.Terminate()

	m := newMixer()
	go m.mix(signalChan)

	// Check GOMAXPROCS
	max_procs := runtime.GOMAXPROCS(-1)
	num_procs_to_set := runtime.NumCPU()
	if max_procs != num_procs_to_set {
		runtime.GOMAXPROCS(num_procs_to_set)
	}

	shellReader := bufio.NewReader(os.Stdin)
	for {
		// SETUP && LOOP
		fmt.Printf(PS2)
		input, err := shellReader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				myexit()
			} else {
				fmt.Println("Got an err for ye: ", err)
			}
		}
		input = strings.TrimSpace(input)

		// ALL DA REGEX FROM HERE ON OUT..

		// SET OFF PRIME TRACK
		prx, _ := regexp.MatchString("^prime$", input)
		if prx {
			go primez(signalChan, tickChan)
		}

		// SET BPM
		bpmre := regexp.MustCompile("^bpm +([0-9]+)$")
		br := bpmre.FindStringSubmatch(input)
		if len(br) == 2 {
			bpmval, err := strconv.ParseFloat(br[1], 64)
			if err != nil {
				fmt.Println("Choked on your Beats Per Minute, mate..")
				continue
			}
			bpm = bpmval
		}

		// CREATE SINE w/ FREQ
		re := regexp.MustCompile("^sine +([0-9]+)$")
		sf := re.FindStringSubmatch(input)
		if len(sf) == 2 {
			freq, err := strconv.ParseFloat(sf[1], 64)
			if err != nil {
				fmt.Println("Choked on your sine freq, mate..")
				continue
			}
			newSine(signalChan, freq)
		}

		// SET SINE ATTRIB
		ssx := regexp.MustCompile("^set sine ([0-9]) ([a-z]+) (0\\.[0-9])$")
		//ssx := regexp.MustCompile("^set sine ([0-9]) ([a-z]+)$")
		//ssx := regexp.MustCompile("^set sine ([0-9]) ([a-z]+)$")
		ssxf := ssx.FindStringSubmatch(input)
		if len(ssxf) == 4 {
			sineNum, err := strconv.Atoi(ssxf[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			val, err := strconv.ParseFloat(ssxf[3], 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if sineNum < len(m.signals) {
				fmt.Println("ooh, got something..", ssxf)
				m.signals[sineNum].amplitude.attack = val

			}
		}
		// 	freq, err := strconv.ParseFloat(sf[1], 64)
		// 	if err != nil {
		// 		fmt.Println("Choked on your sine freq, mate..")
		// 		continue
		// 	}
		// 	newSine(signalChan, freq)
		// }

		// PROCESS LIST
		psx, _ := regexp.MatchString("^ps$", input)
		if psx {
			fmt.Println()
			m.listChans()
			//fmInfo(fmz)
			fmt.Println()
			//sineInfo(sinez)
		}
	}
}
