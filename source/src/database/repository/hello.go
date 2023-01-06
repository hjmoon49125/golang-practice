package repository

import "server_temp/source/src/model"

type helloRepository struct {
}

func HelloRepository() model.HelloRepository {
	return &helloRepository{}
}

func (r *helloRepository) SayHello() string {
	// DB Logic
	return "hello"
}
