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

	sinezChan := make(chan *stereoSine)
	sinez := make([]*stereoSine, 0)

	fmChan := make(chan *FM)
	fmz := make([]*FM, 0)

	tickerChan := make(chan int)
	go ticker(tickerChan)
	go fib(tickerChan)

	portaudio.Initialize()
	defer portaudio.Terminate()

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
			go primez(sinezChan, &sinez)
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

		// CREATE SINE (w/ DEFAULT 440)
		s, _ := regexp.MatchString("^sine$", input)
		if s {
			fmt.Println("Sine o' the times, mate...")
			go newSine(sinezChan, 440) // default
			sinez = append(sinez, <-sinezChan)

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
			go newSine(sinezChan, freq)
			sinez = append(sinez, <-sinezChan)
		}

		// CREATE FREQUENCY MODULATOR
		fmre := regexp.MustCompile("^fm +([0-9]+) +([0-9]+)$")
		fmfz := fmre.FindStringSubmatch(input)
		if len(fmfz) == 3 {
			cfreq, err := strconv.ParseFloat(fmfz[1], 64)
			if err != nil {
				fmt.Println("Choked on your CAR freq, mate..")
				continue
			}
			mfreq, err := strconv.ParseFloat(fmfz[2], 64)
			if err != nil {
				fmt.Println("Choked on your MOD freq, mate..")
				continue
			}
			go newFM(fmChan, cfreq, mfreq)
			fmz = append(fmz, <-fmChan)
		}

		// PROCESS LIST
		psx, _ := regexp.MatchString("^ps$", input)
		if psx {
			fmt.Println()
			fmInfo(fmz)
			fmt.Println()
			sineInfo(sinez)
		}

		// CHANGE A SINE ATTRIBUTE
		c, _ := regexp.MatchString("^set ", input)
		if c {
			setSineProperty(input, sinez)
		}

		// BE QUIET ALL YOU SINEZ
		d, _ := regexp.MatchString("^sssh$", input)
		if d {
			for _, s := range sinez {
				s.set("vol", 0)
			}
		}

		// EVERYBODY STOP!!
		e, _ := regexp.MatchString("^stop$", input)
		if e {
			for _, s := range sinez {
				s.Stop()
				//s.Close()
			}
			sinez = sinez[:0]

			for _, f := range fmz {
				f.Stop()
				//s.Close()
			}
			fmz = fmz[:0]
		}
	}
}

func setSineProperty(inputString string, sinez []*stereoSine) {
	stringieBits := strings.Split(inputString, " ")
	if len(stringieBits) != 4 {
		fmt.Println("Chancer")
		return
	}
	sineToSet, err := strconv.Atoi(stringieBits[1])
	if err != nil {
		fmt.Println("BURNIE CHANCER SINE SELECT ", err)
		return
	}
	propToSet := stringieBits[2]
	valToSet, err := strconv.ParseFloat(stringieBits[3], 64)
	if err != nil {
		fmt.Println("BURNIE CHANCER VAL ", err)
		return
	}
	if sineToSet < len(sinez) {
		fmt.Printf("Changing Sine[%d]\n", sineToSet)
		sinez[sineToSet].set(propToSet, valToSet)
	} else {
		fmt.Println("Chancer")
	}

	//if stringieBits[0] len(sinez)
}

func sineInfo(sinez []*stereoSine) {

	fmt.Println("Sinezzzzzz::")
	for i, d := range sinez {
		fmt.Println("Sine ", i, d)
	}
}

func fmInfo(fmz []*FM) {

	fmt.Println("FMZ::")
	for i, d := range fmz {
		fmt.Println("FM ", i, d)
	}
}
