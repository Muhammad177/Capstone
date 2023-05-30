package request

import (
	"Capstone/models/savedthreads"

	"github.com/go-playground/validator/v10"
)

type Savedthread struct {
	User_id int		`json:"user_id" validate:"required"`
	Thread_id int	`json:"thread_id" validate:"required"`
}

func (req *Savedthread) ToDomain() *savedthreads.Domain {
	return &savedthreads.Domain{
	User_id:	req.User_id,
	Thread_id:	req.Thread_id,
	}
}

func (req *Savedthread) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
