package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"

	"code.google.com/p/portaudio-go/portaudio"
)

var pflag = flag.Bool("primez", false, "Play Primez rather than interactive")

func main() {
	flag.Parse()
	//cmds := []string{"ls", "exit", "jobbie"}

	PS2 := ansi.Color("#CMDSine> ", "magenta")

	signalChan := make(chan SoundGen)

	tickChan := make(chan int)
	go ticker(tickChan)
	// go fib(tickerChan)

	portaudio.Initialize()
	defer portaudio.Terminate()

	m := newMixer()
	go m.mix(signalChan)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Check GOMAXPROCS
	max_procs := runtime.GOMAXPROCS(-1)
	num_procs_to_set := runtime.NumCPU()
	if max_procs != num_procs_to_set {
		runtime.GOMAXPROCS(num_procs_to_set)
	}

	// non-interactive - running Prime loop
	if *pflag == true {
		forevs := make(chan bool)
		go primez(signalChan, tickChan)
		<-forevs
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
			tickLength = 60000 / bpm / 60
			loopLength = tickLength * 16
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
		// ssx := regexp.MustCompile("^set sine ([0-9]) ([a-z]+) ([0-9\\.]+)$")
		// ssxf := ssx.FindStringSubmatch(input)
		// if len(ssxf) == 4 {
		// 	sineNum, err := strconv.Atoi(ssxf[1])
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		continue
		// 	}
		// 	attrib := ssxf[2]
		// 	val, err := strconv.ParseFloat(ssxf[3], 64)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		continue
		// 	}
		// 	if sineNum < len(m.signals) {
		// 		m.signals[sineNum].set(attrib, val)
		// 	}
		// }

		// CREATE FREQMOD
		fmx := regexp.MustCompile("^fm +([0-9]+) +([0-9]+)$")
		fmxr := fmx.FindStringSubmatch(input)
		// if len(fmxr) == 2 {
		if len(fmxr) > 1 {
			fmt.Println("New FM!")
			car, err := strconv.ParseFloat(fmxr[1], 64)
			if err != nil {
				fmt.Println("Choked on your carrier freq, mate..")
				continue
			}
			mod, err := strconv.ParseFloat(fmxr[2], 64)
			if err != nil {
				fmt.Println("Choked on your modulator freq, mate..")
				continue
			}
			newFM(signalChan, car, mod)
		}

		// PROCESS LIST
		psx, _ := regexp.MatchString("^ps$", input)
		if psx {
			fmt.Printf("///CMDSine:: bpm: %.0f \\\\\\\n", bpm)
			m.listChans()
			fmt.Println()
		}
	}
}
