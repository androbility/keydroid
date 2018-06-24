package main

import (
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/eiannone/keyboard"
)

func main() {
	cmndr, err := New()
	if err != nil {
		log.Fatalf("error connecting to adb server: %s\n\nDid you run adb connect?\n", err)
	}

	for {
		ch, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		// Ctrl key values return [0, uint16], so convert the uint16
		// to a rune.
		if ch == 0 {
			ch = rune(key)
		}

		// Quit on q, Q, and Ctrl-C.
		if (unicode.ToLower(ch) == 'q') || (ch == '\x03') {
			cmndr.Quit()
		}

		cmndr.Write(Keycode(ch))
	}
}
