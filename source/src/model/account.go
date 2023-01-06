package model

type AccountReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Account struct {
	Account  string `gorm:"account"`
	Password string `gorm:"password"`
}

type AccountRepository interface {
	FindByAccount(acc string) (*Account, error)
}

type AccountService interface {
	Login(acc string, pass string) (*Account, error)
}