package xild

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/caarlos0/env"

	"os"
)

type GELFFormatter struct {
}

type Message struct {
	Version   string                 `json:"version"`
	Host      string                 `json:"host"`
	Short     string                 `json:"short_message"`
	Full      string                 `json:"full_message"`
	Timestamp string                 `json:"timestamp"`
	Level     int32                  `json:"level"`
	File      string                 `json:"file"`
	Line      int                    `json:"line"`
	Extra     map[string]interface{} `json:"-"`
}

type ExtraEnvField struct {
	Facility string `env:"APPLICATION_NAME"`
}

var cfg ExtraEnvField

func init() {
	cfg := ExtraEnvField{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", cfg)
}

func (f *GELFFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	host, err := os.Hostname()
	if err != nil {
		host = "localhost"
	}

	// remove trailing and leading whitespace
	p := bytes.TrimSpace([]byte(entry.Message))

	// If there are newlines in the message, use the first line
	// for the short message and set the full message to the
	// original input.  If the input has no newlines, stick the
	// whole thing in Short.
	short := p
	full := []byte("")
	if i := bytes.IndexRune(p, '\n'); i > 0 {
		short = p[:i]
		full = p
	}

	level := int32(entry.Level) + 2 // logrus levels are lower than syslog by 2

	extra := map[string]interface{}{}
	//set facility as application name or null
	extra["_facility"] = cfg.Facility
	// Merge extra fields
	for f, v := range entry.Data {
		f = fmt.Sprintf("_%s", f) // "[...] every field you send and prefix with a _ (underscore) will be treated as an additional field."
		extra[f] = v
	}

	m := Message{
		Version:   "1.1",
		Host:      host,
		Short:     string(short),
		Full:      string(full),
		Timestamp: entry.Time.Format(logrus.DefaultTimestampFormat),
		Level:     level,
		// File:     entry.file, deprecated
		// Line:     entry.line, deprecated
		Extra: extra,
	}

	serialized, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
