package service

import (
	"github.com/BaytoorJr/sso/src/repository"
	"github.com/go-kit/kit/log"
)

type service struct {
	store  repository.MainRepo
	logger log.Logger
}

func NewService(store repository.MainRepo, logger log.Logger) MainService {
	return &service{
		store:  store,
		logger: logger,
	}
}
