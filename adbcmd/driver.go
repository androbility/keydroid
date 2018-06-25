package adbcmd

import (
	"errors"
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
	ch         chan time.Time
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

	cmndr := &Commander{
		cmd:        cmd,
		in:         stdin,
		lastActive: time.Now(),
		m:          &sync.Mutex{},
		ch:         make(chan time.Time),
	}
	go cmndr.ping(149 * time.Second)

	return cmndr, nil
}

func (c *Commander) Write(key Keycode) error {
	code, ok := key.Event()
	if !ok {
		log.WithFields(log.Fields{
			"key": key.Rune(),
		}).Error("Key not bound")

		return nil
	}

	if _, err := c.in.Write(code); err != nil {
		// Communication with the Android device failed.
		log.WithFields(log.Fields{
			"error": err,
			"key":   rune(key),
		}).Error("KeyEvent send failed")

		// We can assume the server is down, or restarting.
		// Let's close the channel, kill cmd, and quit.
		defer close(c.ch)
		defer c.cmd.Wait()

		return errors.New("server connection lost")
	}
	go c.touch()

	log.Info(strings.Trim(string(code), "\n"))

	return nil
}

func (c *Commander) Quit() {
	c.in.Write([]byte("exit\n"))
	log.Info("Quitting")
	c.cmd.Wait()
	os.Exit(0)
}

// Ensure device stays awake.  Purpose: FireTV Cube.
func (c *Commander) ping(dur time.Duration) {
	for range c.ch {
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

// Wait for server to reconnect.
func WaitForAndroid() {
	log.Info("Waiting for a new adb connection.  (Hint: adb connect <ip[:port]>)")
	exec.Command("adb", "wait-for-device").Run()
	log.Info("Success!")
}
