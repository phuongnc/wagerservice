package logger

import (
	"path/filepath"
	"time"
	"wagerservice/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"

	"os"
)

func Init(cfg *config.Config) (*logrus.Logger, error) {
	configLog := cfg.LogConfig
	log := logrus.New()
	switch configLog.Format {
	case "json":
		log.SetFormatter(new(logrus.JSONFormatter))
	default:
		log.SetFormatter(new(logrus.TextFormatter))
	}
	if configLog.Output != "" {
		switch configLog.Output {
		case "stdout":
			log.SetOutput(os.Stdout)
		case "stderr":
			log.SetOutput(os.Stderr)
		case "file":
			wd, _ := os.Getwd()
			dir := filepath.Join(wd, configLog.Path)
			_ = os.Mkdir(dir, 0777)
			linkName := filepath.Join(dir, "wagerservice")
			write, err := rotatelogs.New(
				filepath.Join(dir, "wagerservice-%Y-%m-%d.log"),
				rotatelogs.WithMaxAge(time.Duration(24*configLog.Expired)*time.Hour),
				rotatelogs.WithLinkName(linkName),
				rotatelogs.WithRotationTime(24*time.Hour),
			)
			if err != nil {
				return nil, err
			}
			log.SetOutput(write)
		}
	}
	return log, nil
}
