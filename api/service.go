//
//  Practicing Redis
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Service represent the services
type Service interface {
	Set(key string, value interface{}, exp time.Duration) (int, error)
	SAdd(key string, value ...interface{}) (int, error)
	HSet(key, field, value string) (bool, int, error)
	Get(key string) (string, int, error)
	GetSAdd(key string) ([]string, int, error)
	GetHSet(key, field string) (string, int, error)
	Delete(key string) (int, error)
}

type implService struct {
	log        *logrus.Entry
	repository Repository
}

// NewService will create an object that represent the Service interface
func NewService(log *logrus.Entry, r Repository) Service {
	return &implService{log: log, repository: r}
}

func (s *implService) Set(key string, value interface{}, exp time.Duration) (int, error) {
	err := s.repository.Set(key, value, exp)
	if err != nil {
		s.log.Errorf("can't create set data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}

func (s *implService) SAdd(key string, value ...interface{}) (int, error) {
	err := s.repository.SAdd(key, value)
	if err != nil {
		s.log.Errorf("can't create sadd data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}

func (s *implService) HSet(key, field, value string) (bool, int, error) {
	isCreate, err := s.repository.HSet(key, field, value)
	if err != nil {
		s.log.Errorf("can't create hset data: %s", err.Error())
		return isCreate, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return isCreate, 0, nil
}

func (s *implService) Get(key string) (string, int, error) {
	data, err := s.repository.Get(key)
	if err != nil {
		s.log.Errorf("can't get set data: %s", err.Error())
		return "", http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return data, 0, nil
}

func (s *implService) GetSAdd(key string) ([]string, int, error) {
	datas, err := s.repository.GetSadd(key)
	if err != nil {
		s.log.Errorf("can't get sadd data: %s", err.Error())
		return nil, http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return datas, 0, nil
}

func (s *implService) GetHSet(key, field string) (string, int, error) {
	data, err := s.repository.GetHSet(key, field)
	if err != nil {
		s.log.Errorf("can't get hset data: %s", err.Error())
		return "", http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return data, 0, nil
}

func (s *implService) Delete(key string) (int, error) {
	err := s.repository.Delete(key)
	if err != nil {
		s.log.Errorf("can't delete data: %s", err.Error())
		return http.StatusInternalServerError, errors.New("Oops! Something went wrong. Please try again later")
	}

	return 0, nil
}