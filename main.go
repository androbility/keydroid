package main

import (
	"fmt"
	"unicode"

	log "github.com/Sirupsen/logrus"
	"github.com/androbility/adbi"
	"github.com/eiannone/keyboard"
)

func main() {
	bindings := adbi.LoadConfigFile("$HOME/.keydroid", defaultBindings)
	for {
		adbi.WaitForAndroid()
		err := Watch(bindings)
		log.Error(err)
	}
}

func Watch(keymap map[rune]adbi.Keyevent) error {
	cmndr, err := adbi.New()
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

		// We need to lookup the key.  Is it defined?
		event, ok := keymap[ch]
		if !ok {
			continue
		}

		// If KEYCODE_UNKNOWN, we need to send Raw instead of Signal.
		if event == adbi.KEYCODE_UNKNOWN {
			switch ch {
			case 117:
				cmndr.Raw("swipe 700 120 700 0")
			case 100:
				cmndr.Raw("swipe 700 0 700 120")
			}

			continue
		}

		if err = cmndr.Signal(event); err != nil {
			return err
		}
	}
}

// Default keybindings for keydroid.
var defaultBindings = `{
	"keybindings": {
		"\b":   "KEYCODE_BACK",
		"\t":   "KEYCODE_TAB",
		"\r":   "KEYCODE_ENTER",
		" ":    "KEYCODE_SPACE",
		"-":    "KEYCODE_VOLUME_DOWN",
		"=":    "KEYCODE_VOLUME_UP",
		"+":    "KEYCODE_VOLUME_UP",
		"1":    "KEYCODE_TV_INPUT_HDMI_1",
		"2":    "KEYCODE_TV_INPUT_HDMI_2",
		"3":    "KEYCODE_TV_INPUT_HDMI_3",
		"4":    "KEYCODE_TV_INPUT_HDMI_4",
		"a":    "KEYCODE_ASSIST",
		"b":    "KEYCODE_BACK",
		"c":    "KEYCODE_DPAD_CENTER",
		"f":    "KEYCODE_MEDIA_FAST_FORWARD",
		"h":    "KEYCODE_DPAD_LEFT",
		"i":    "KEYCODE_TV_INPUT",
		"j":    "KEYCODE_DPAD_DOWN",
		"k":    "KEYCODE_DPAD_UP",
		"l":    "KEYCODE_DPAD_RIGHT",
		"m":    "KEYCODE_MUTE",
		"o":    "KEYCODE_HOME",
        "p":    "KEYCODE_MEDIA_PLAY_PAUSE",
		"r":    "KEYCODE_MEDIA_REWIND",
		"s":    "KEYCODE_MEDIA_STOP",
		"S":    "KEYCODE_SLEEP",
		"t":    "KEYCODE_TV_TIMER_PROGRAMMING",
		"u":    "KEYCODE_MENU",
		"v":    "KEYCODE_HELP",
		"w":    "KEYCODE_WAKEUP"
	},
    "rawbindings": {
		"U":    "swipe 700 120 700 0",
		"D":    "swipe 700 0 700 120"
    }
}
`
