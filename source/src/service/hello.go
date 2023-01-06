package service

import (
	"server_temp/source/src/database/repository"
	"server_temp/source/src/model"
)

type helloService struct {
	helloRepo model.HelloRepository
}

func HelloService() model.HelloService {
	return &helloService{
		helloRepo: repository.HelloRepository(),
	}
}

func (s *helloService) SayHello() string {
	return s.helloRepo.SayHello()
}