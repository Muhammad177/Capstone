package request

import (
	"Capstone/models/likes"

	"github.com/go-playground/validator/v10"
)

type Like struct {
	User_id int		`json:"user_id" validate:"required"`
	Thread_id int	`json:"thread_id" validate:"required"`
}

func (req *Like) ToDomain() *likes.Domain {
	return &likes.Domain{
	User_id:	req.User_id,
	Thread_id:	req.Thread_id,
	}
}

func (req *Like) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
