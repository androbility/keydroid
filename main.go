package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/eiannone/keyboard"
	adb "github.com/yosemite-open/go-adb"
)

func main() {
	_, err := adb.New()
	if err != nil {
		log.Fatalf("error connecting to adb server: %s\n\nDid you run adb connect?\n", err)
	}

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()
	stdin := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		switch ch {
		case 'H', 'h':
			log.Info("Move left")
		case 'J', 'j':
			log.Info("Move up")
		case 'K', 'k':
			log.Info("Move down")
		case 'L', 'l':
			log.Info("Move right")
		case 'I', 'i':
			log.Info("Begin input")
			fmt.Print("> ")
			stdin.Scan()
			log.Infof("You entered: %s", stdin.Text())
		case 'Q', 'q':
			log.Info("Quitting")
			os.Exit(0)
		}
	}
}
