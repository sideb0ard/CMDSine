package main

import (
	"bufio"
	"fmt"
	"os"
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
		switch input {
		case "sine":
			fmt.Println("Sine o' the times, mate...")
			go newSine(sinezChan)
			sinez = append(sinez, <-sinezChan)
		case "sines":
			fmt.Println("Sinezzzzzz", sinez)
			for i := range sinez {
				fmt.Println("Sine ", i)
			}
		case "exit":
			myexit()
		case "help":
			help()
		case "jobbie":
			fmt.Println("smellybum")
		case "":
			continue
		default:
			fmt.Println("I DONT UNDERSTAND YOU HUMAN")
			help()
		}
	}
}
