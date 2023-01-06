package repository

import (
	"fmt"
	"server_temp/source/src/model"
)

type accountRepository struct {
}

func AccountRepository() model.AccountRepository {
	return &accountRepository{}
}

func (r *accountRepository) FindByAccount(acc string) (*model.Account, error) {
	// DB Logic
	account := model.Account{
		Account: "Yaaaaman",
	}

	if acc != account.Account {
		return nil, fmt.Errorf("account not found")
	}

	return &account, nil
}
