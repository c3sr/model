package model

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	log      *logrus.Entry
	validate = validator.New()
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "model")
	})
}
