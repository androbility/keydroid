package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/eiannone/keyboard"
	"github.com/stephenwithav/keydroid/shell"
)

func main() {
	cmd, stdin, stdout, err := shell.New()
	if err != nil {
		log.Fatalf("error connecting to adb server: %s\n\nDid you run adb connect?\n", err)
	}
	defer stdin.Close()
	defer stdout.Close()

	// For keyboard input
	scanIn := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		switch ch {
		case 'H', 'h':
			log.Info("Move left")
			_, err := stdin.Write([]byte("input keyevent KEYCODE_DPAD_LEFT\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'J', 'j':
			log.Info("Move down")
			_, err := stdin.Write([]byte("input keyevent KEYCODE_DPAD_DOWN\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'K', 'k':
			log.Info("Move down")
			_, err := stdin.Write([]byte("input keyevent KEYCODE_DPAD_UP\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'L', 'l':
			log.Info("Move right")
			_, err := stdin.Write([]byte("input keyevent KEYCODE_DPAD_RIGHT\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'I', 'i':
			log.Info("Begin input")
			fmt.Print("> ")
			scanIn.Scan()
			log.Infof("You entered: %s", scanIn.Text())
			_, err := stdin.Write([]byte(scanIn.Text()))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'X', 'x':
			log.Info("Pressed Enter")
			_, err := stdin.Write([]byte("input keyevent KEYCODE_ENTER\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
		case 'Q', 'q':
			log.Info("Quitting")
			_, err := stdin.Write([]byte("exit\n"))
			if err != nil {
				log.Errorf("err: %s")
			}
			cmd.Wait()
			os.Exit(0)
		}
	}
}
