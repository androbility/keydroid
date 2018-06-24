package main

import (
	"fmt"
	"unicode"
)

type Keycode rune

func (k Keycode) Event() ([]byte, bool) {
	if code, ok := keymap[unicode.ToLower(rune(k))]; ok {
		return []byte(fmt.Sprintf("input keyevent %s\n", code)), true
	}

	return nil, false
}

func (k Keycode) Rune() rune {
	return rune(k)
}

var (
	keymap = map[rune]string{
		' ': "KEYCODE_SPACE",
		'-': "KEYCODE_VOLUME_DOWN",
		'=': "KEYCODE_VOLUME_UP",
		'+': "KEYCODE_VOLUME_UP",
		'1': "KEYCODE_TV_INPUT_HDMI_1",
		'2': "KEYCODE_TV_INPUT_HDMI_2",
		'3': "KEYCODE_TV_INPUT_HDMI_3",
		'4': "KEYCODE_TV_INPUT_HDMI_4",
		'a': "KEYCODE_VOICE_ASSIST",
		'b': "KEYCODE_BACK",
		'f': "KEYCODE_MEDIA_FAST_FORWARD",
		'h': "KEYCODE_DPAD_LEFT",
		'j': "KEYCODE_DPAD_DOWN",
		'k': "KEYCODE_DPAD_UP",
		'l': "KEYCODE_DPAD_RIGHT",
		'm': "KEYCODE_VOLUME_MUTE",
		'r': "KEYCODE_MEDIA_REWIND",
		's': "KEYCODE_MEDIA_STOP",
		't': "KEYCODE_TV_TIMER_PROGRAMMING",
		'v': "KEYCODE_TV_POWER",
		'w': "KEYCODE_WAKEUP",
	}
)
