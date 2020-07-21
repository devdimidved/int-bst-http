package bst

import (
	"github.com/sirupsen/logrus"
	"time"
)

type loggingService struct {
	logger *logrus.Logger
	Service
}

func NewLoggingService(logger *logrus.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) Insert(v int) {
	defer func(begin time.Time) {
		logger := s.logger.WithFields(logrus.Fields{
			"method": "insert",
			"val":    v,
		})
		logger.Infof("took %v", time.Since(begin))
	}(time.Now())

	s.Service.Insert(v)
}

func (s *loggingService) Search(v int) bool {
	defer func(begin time.Time) {
		logger := s.logger.WithFields(logrus.Fields{
			"method": "search",
			"val":    v,
		})
		logger.Infof("took %v", time.Since(begin))
	}(time.Now())

	return s.Service.Search(v)
}

func (s *loggingService) Remove(v int) {
	defer func(begin time.Time) {
		logger := s.logger.WithFields(logrus.Fields{
			"method": "remove",
			"val":    v,
		})
		logger.Infof("took %v", time.Since(begin))
	}(time.Now())

	s.Service.Remove(v)
}
