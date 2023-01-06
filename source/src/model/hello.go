package model

type Hello struct {
	Message string `json:"message"`
}

type HelloRepository interface {
	SayHello() string
}

type HelloService interface {
	SayHello() string
}