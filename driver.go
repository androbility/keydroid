package main

import (
	"io"
	"os"
	"os/exec"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type Commander struct {
	cmd *exec.Cmd
	in  io.WriteCloser
}

func New() (*Commander, error) {
	cmd := exec.Command("adb", "shell")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &Commander{cmd, stdin}, nil
}

func (c *Commander) Write(key Keycode) {
	if code, ok := key.Event(); ok {
		if _, err := c.in.Write(code); err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"key":   rune(key),
			}).Error("KeyEvent send failed")
		}

		log.Info(strings.Trim(string(code), "\n"))

		return
	}

	log.WithFields(log.Fields{
		"key": key.Rune(),
	}).Error("Key not bound")
}

func (c *Commander) Exec(cmd string) error {
	_, err := c.in.Write([]byte(cmd))
	return err
}

func (c *Commander) Quit() {
	c.in.Write([]byte("exit\n"))
	log.Info("Quitting")
	c.cmd.Wait()
	os.Exit(0)
}
