package config

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	log.SetLevel(GetLoggerLevel())
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TrimMessages:    true,
		ShowFullLevel:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func GetLoggerLevel() log.Level {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		return log.InfoLevel
	}

	switch level {
	case "DEBUG":
		return log.DebugLevel
	case "INFO":
		return log.InfoLevel
	case "WARN":
		return log.WarnLevel
	case "ERROR":
		return log.ErrorLevel
	case "FATAL":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}