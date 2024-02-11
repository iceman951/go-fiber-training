package config

import (
	"fmt"
	"os"
	"time"

	"github.com/baac-tech/zlogwrap"
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const (
	TimeFormat = time.RFC3339Nano
)

func Setup(configPath string) {
	if os.Getenv("ENV") == "" {
		panic(fmt.Errorf("📌 PLEASE SET 'ENV' ex. 'export ENV=dev' 📌"))
	}
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "config.Setup",
	})
	viper.SetConfigType("json")
	viper.SetConfigName("env." + os.Getenv("ENV"))
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s %v", err, "\n"))
	}
	viper.Debug()

	SetupLogLevel()

	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		logger.Info("File Edited:", event.Name)
		SetupLogLevel()
	})
}

func SetupLogLevel() {
	logger := zlogwrap.New(zlogwrap.Config{
		ServiceName: "Service.SetupLogLevel",
	})

	logLevel := viper.GetString("LogLevel")
	zerolog.SetGlobalLevel(GetLogLevel(logLevel))

	logger.Info("LogLevel Changed:", logLevel)
}

func GetLogLevel(level string) zerolog.Level {
	switch level {
	case zerolog.LevelPanicValue: // 5
		return zerolog.PanicLevel
	case zerolog.LevelFatalValue: // 4
		return zerolog.FatalLevel
	case zerolog.LevelErrorValue: // 3
		return zerolog.ErrorLevel
	case zerolog.LevelWarnValue: // 2
		return zerolog.WarnLevel
	case zerolog.LevelInfoValue: // 1
		return zerolog.InfoLevel
	case zerolog.LevelDebugValue: // 0
		return zerolog.DebugLevel
	case zerolog.LevelTraceValue: // -1
		return zerolog.TraceLevel
	default:
		return zerolog.DebugLevel // Debug is default
	}
}
