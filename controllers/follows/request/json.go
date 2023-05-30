package request

import (
	"Capstone/models/follows"

	"github.com/go-playground/validator/v10"
)

type Follow struct {
	User_id int		`json:"user_id" validate:"required"`
	Thread_id int	`json:"thread_id" validate:"required"`
}

func (req *Follow) ToDomain() *follows.Domain {
	return &follows.Domain{
	User_id:	req.User_id,
	Thread_id:	req.Thread_id,
	}
}

func (req *Follow) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
