package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	portaudio.Initialize()
	defer portaudio.Terminate()

	shellReader := bufio.NewReader(os.Stdin)
	for {
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

		s, _ := regexp.MatchString("^sine$", input)
		if s {
			fmt.Println("Sine o' the times, mate...")
			go newSine(sinezChan, 440)
			sinez = append(sinez, <-sinezChan)
		}
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

		b, _ := regexp.MatchString("^sines$", input)
		if b {
			sineInfo(sinez)
		}

		fmre := regexp.MustCompile("^fm +([0-9]+) +([0-9]+)$")
		fmfz := fmre.FindStringSubmatch(input)
		fmt.Println("GOt THIS: ", len(fmfz))
		if len(fmfz) == 3 {
			mfreq, err := strconv.ParseFloat(fmfz[1], 64)
			if err != nil {
				fmt.Println("Choked on your MOD freq, mate..")
				continue
			}
			cfreq, err := strconv.ParseFloat(fmfz[2], 64)
			if err != nil {
				fmt.Println("Choked on your CAR freq, mate..")
				continue
			}
			go newFM(cfreq, mfreq)
			//sinez = append(sinez, <-sinezChan)
		}
		c, _ := regexp.MatchString("^set ", input)
		if c {
			setSineProperty(input, sinez)
		}

		d, _ := regexp.MatchString("^sssh$", input)
		if d {
			for _, s := range sinez {
				s.set("vol", 0)
			}
		}

		e, _ := regexp.MatchString("^stop$", input)
		if e {
			for _, s := range sinez {
				s.Stop()
				//s.Close()
			}
			sinez = sinez[:0]
		}
		//case input.MatchString("sine"):
		//case "sines":
		//case input.MatchString("sine"):
		//case "set *":
		//	setSineProperty(input)
		//case "exit":
		//	myexit()
		//case "help":
		//	help()
		//case "jobbie":
		//	fmt.Println("smellybum")
		//case "":
		//	continue
		//default:
		//	fmt.Println("I DONT UNDERSTAND YOU HUMAN")
		//	help()
		//}
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
