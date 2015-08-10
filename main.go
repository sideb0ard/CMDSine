package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"code.google.com/p/portaudio-go/portaudio"
)

func main() {
	//cmds := []string{"ls", "exit", "jobbie"}

	sinezChan := make(chan *stereoSine)
	sinez := make([]*stereoSine, 0)

	portaudio.Initialize()
	defer portaudio.Terminate()

	shellReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("<CMDzine#> ")
		input, err := shellReader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				myexit()
			} else {
				fmt.Println("Got an err for ye: ", err)
			}
		}
		input = strings.TrimSpace(input)

		a, _ := regexp.MatchString("^sine$", input)
		if a {
			fmt.Println("Sine o' the times, mate...")
			go newSine(sinezChan)
			sinez = append(sinez, <-sinezChan)
		}
		b, _ := regexp.MatchString("^sines$", input)
		if b {
			sineInfo(sinez)
		}

		c, _ := regexp.MatchString("^set ", input)
		if c {
			setSineProperty(input, sinez)
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
	fmt.Println("BAERF - SET SINE PROPEZ", inputString, sinez)
	stringieBits := strings.Split(inputString, " ")
	fmt.Println("BAERF - SET SINE PROPEZ", inputString, stringieBits)
	fmt.Println("LEN SINEZ ", len(stringieBits))
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
		fmt.Printf("Changing Sine[%d] - %s -- %d\n", sineToSet, propToSet, valToSet)
		sinez[sineToSet].set(propToSet, valToSet)
	} else {
		fmt.Println("Chancer")
	}

	//if stringieBits[0] len(sinez)
}

func sineInfo(sinez []*stereoSine) {

	fmt.Println("Sinezzzzzz", sinez)
	for i, d := range sinez {
		fmt.Println("Sine ", i, d)
	}
}
