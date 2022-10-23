package registry

import (
	"wagerservice/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Provider struct {
	DB     *gorm.DB
	Config *config.Config
	Logger *logrus.Logger
}
