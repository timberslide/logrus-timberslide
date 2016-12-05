package tslogrus

import (
	"github.com/Sirupsen/logrus"

	"github.com/timberslide/gotimberslide"
)

const (
	// DefaultHost is used when the client does not supply a host
	DefaultHost = "gw.timberslide.com:443"
)

// Hook contains a Timberslide hook
type Hook struct {
	Token string // Token is the client's Timberslide token
	Host  string // Host is optional, and does not usually need to be set
	Topic string // Topic is the Timberslide topic name to use

	channel ts.Channel
	client  ts.Client
}

// NewTimberslideHook creates a new Logrus hook for logging to Timberslide
func NewTimberslideHook(hook *Hook) (*Hook, error) {
	var err error
	if hook.Host == "" {
		hook.Host = DefaultHost
	}
	hook.client, err = ts.NewClient(hook.Host, hook.Token)
	if err != nil {
		return hook, err
	}
	err = hook.client.Connect()
	if err != nil {
		return hook, err
	}
	hook.channel, err = hook.client.CreateChannel(hook.Topic)
	return hook, err
}

// Fire sends an event to Timberslide
func (h *Hook) Fire(entry *logrus.Entry) error {
	s, err := entry.String()
	if err != nil {
		return err
	}
	err = h.channel.Send(s) // This should be all the goodies
	return err
}

// Levels returns all available log levels
func (h *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
