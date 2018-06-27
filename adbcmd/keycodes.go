package adbcmd

import (
	"fmt"
)

type Keycode uint32

func (k Keycode) Event() ([]byte, bool) {
	if code, ok := keymap[rune(k)]; ok {
		return []byte(fmt.Sprintf("input keyevent %d\n", code)), true
	}

	return nil, false
}

func (k Keycode) Rune() rune {
	return rune(k)
}
