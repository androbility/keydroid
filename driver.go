package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Commander struct {
	cmd        *exec.Cmd
	in         io.WriteCloser
	lastActive time.Time
	m          *sync.Mutex
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

	cmndr := &Commander{cmd, stdin, time.Now(), &sync.Mutex{}}
	go cmndr.keepAwake(149 * time.Second)

	return cmndr, nil
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

// Ensure device stays awake.  Purpose: FireTV Cube.
func (c *Commander) keepAwake(dur time.Duration) {
	ch := time.Tick(dur)

	for range ch {
		c.touch()
	}
}

func (c *Commander) touch() {
	if now := time.Now(); now.Sub(c.lastActive) > time.Minute {
		c.m.Lock()
		defer c.m.Unlock()

		c.Write(Keycode('w'))
		c.lastActive = now
	}
}
