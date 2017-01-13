package logrus3164

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
)

const format = "Jan 2 15:04:05"

type hookImpl struct {
	io.Writer
	hostname string
	tag      string
}

func NewHook(writer io.Writer, tag string) (logrus.Hook, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return &hookImpl{writer, hostname, tag}, nil
}

func (*hookImpl) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *hookImpl) Fire(entry *logrus.Entry) error {
	date := time.Now().Format(format)
	msg, err := entry.String()
	if err != nil {
		return err
	}
	payload := fmt.Sprintf(
		"<22>%s %s %s: %s",
		date,
		hook.hostname,
		hook.tag,
		msg,
	)
	_, err = hook.Write([]byte(payload))
	return err
}
