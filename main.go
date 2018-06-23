package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/eiannone/keyboard"
)

func main() {
	cmndr, err := New()
	if err != nil {
		log.Fatalf("error connecting to adb server: %s\n\nDid you run adb connect?\n", err)
	}

	// For keyboard input
	stdin := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if unicode.IsSpace(ch) {
			ch = '\n'
		}

		switch ch {
		case 'I', 'i':
			log.Info("Begin input")
			fmt.Print("> ")
			stdin.Scan()
			log.Infof("Sending: %s", stdin.Text())
			err := cmndr.Exec(stdin.Text())
			if err != nil {
				log.Errorf("err: %s")
			}
		case 0:
			log.Info("Pressed Enter")
			if err := cmndr.Exec("input keyevent KEYCODE_ENTER\n"); err != nil {
				log.Errorf("err: %s")
			}
		case 'Q', 'q':
			cmndr.Quit()
		default:
			cmndr.Write(Keycode(ch))
		}
	}
}
