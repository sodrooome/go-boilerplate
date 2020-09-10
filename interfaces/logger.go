package interfaces

import (
	"fmt"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type LoggerInterface interface {
	Error(args ...interface{})
}

//goland:noinspection ALL
func CreateLogger(client *raven.Client) LoggerInterface {
	logger := logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: true},
		Hooks: make(logrus.LevelHooks),
		Level: logrus.InfoLevel,
	}

	if client != nil {
		// integrate logrus with sentry
		hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		// WIP: create a function for running with logrus + sentry
		timeout := viper.GetInt("SENTRY_TIMEOUT")
		hook.Timeout = time.Duration(timeout) * time.Second
		hook.StacktraceConfiguration.Enable = true
		if err != nil {
			fmt.Println("Getting an error")
		}
	}
	return &logger
}