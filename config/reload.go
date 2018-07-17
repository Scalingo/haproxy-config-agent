package config

import (
	"io/ioutil"
	"os"
	"strconv"
)

type Reloader interface {
	Reload() error
}

type ConfigReloader struct {
	pidFile string
	signal  os.Signal
}

func NewConfigReloader(pidFile string, signal os.Signal) *ConfigReloader {
	return &ConfigReloader{
		pidFile: pidFile,
		signal:  signal,
	}
}

func (c *ConfigReloader) Reload() error {
	pidBytes, err := ioutil.ReadFile(c.pidFile)
	if err != nil {
		return err
	}

	pid, err := strconv.Atoi(string(pidBytes))
	if err != nil {
		return err
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Signal(c.signal)
}
