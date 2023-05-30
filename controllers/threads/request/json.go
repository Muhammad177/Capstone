package request

import (
	"Capstone/models/threads"

	"github.com/go-playground/validator/v10"
)

type Thread struct {
	Title	string	`json:"tittle" validate:"required"`
	Topic   string	`json:"topic" validate:"required"`
	Content string 	`json:"content" validate:"required"`
	File    string	`json:"file" validate:"required"`
	User_id int		`json:"user_id" validate:"required"`
	Role	string	`json:"role" validate:"required"`
}

func (req *Thread) ToDomain() *threads.Domain {
	return &threads.Domain{
	Title:	req.Title,
	Topic:	req.Topic,
	Content: req.Content,
	File:    req.File,
	User_id:	req.User_id,
	Role:	req.Role,
	}
}

func (req *Thread) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
