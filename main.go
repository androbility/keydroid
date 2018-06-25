package main

import (
	"fmt"
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/eiannone/keyboard"
	"github.com/stephenwithav/keydroid/adbcmd"
)

func main() {
	for {
		adbcmd.WaitForAndroid()
		err := Watch()
		log.Error(err)
	}
}

func Watch() error {
	cmndr, err := adbcmd.New()
	if err != nil {
		return fmt.Errorf("error connecting to adb server: %s", err)
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

		if err = cmndr.Write(adbcmd.Keycode(ch)); err != nil {
			return err
		}
	}
}
