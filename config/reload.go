package config

import "os"

type Reloader interface {
	Reload() error
}

type ConfigReloader struct {
	pid    int
	signal os.Signal
}

func NewConfigReloader(pid int, signal os.Signal) *ConfigReloader {
	return &ConfigReloader{
		pid:    pid,
		signal: signal,
	}
}

func (c *ConfigReloader) Reload() error {
	process, err := os.FindProcess(c.pid)
	if err != nil {
		return err
	}

	return process.Signal(c.signal)
}
