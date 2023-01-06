package service

import (
	// "fmt"
	"server_temp/source/src/database/repository"
	"server_temp/source/src/model"
	// "server_temp/source/src/utils"
)

type accountService struct {
	accRepo model.AccountRepository
}

func AccountService() model.AccountService {
	return &accountService{
		accRepo: repository.AccountRepository(),
	}
}

func (s *accountService) Login(acc string, pass string) (*model.Account, error) {
	account, err := s.accRepo.FindByAccount(acc)
	if err != nil {
		return nil, err
	}

	// if !utils.CompareSomeShit(account.Password, pass) {
	// 	return nil, fmt.Errorf("invalid password")
	// }

	return account, err
}