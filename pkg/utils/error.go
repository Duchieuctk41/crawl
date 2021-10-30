package utils

import "github.com/sirupsen/logrus"

func LogError(err error) {
	if err != nil {
		logrus.Info(err)
	}
}
