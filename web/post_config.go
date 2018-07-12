package web

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Scalingo/haproxy-config-agent/config"
)

type ConfigManager struct {
	Reloader config.Reloader
	Writer   io.Writer
}

func (c ConfigManager) PostConfig(w http.ResponseWriter, r *http.Request, params map[string]string) error {
	if r.Method != "POST" {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
		return nil
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	_, err = c.Writer.Write(body)
	if err != nil {
		return err
	}

	err = c.Reloader.Reload()
	if err != nil {
		return err
	}

	return nil
}
