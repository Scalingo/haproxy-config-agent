package config

import "io/ioutil"

type ConfigWriter struct {
	path string
}

func NewConfigWriter(path string) *ConfigWriter {
	return &ConfigWriter{
		path: path,
	}
}

func (c *ConfigWriter) Write(payload []byte) (int, error) {
	err := ioutil.WriteFile(c.path, payload, 0644)
	if err != nil {
		return 0, err
	}

	return len(payload), nil
}
