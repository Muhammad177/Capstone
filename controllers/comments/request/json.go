package request

import (
	"Capstone/models/comments"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	User_id int		`json:"user_id" validate:"required"`
	Thread_id int	`json:"thread_id" validate:"required"`
	Comment	string	`json:"comment" validate:"required"`
}

func (req *Comment) ToDomain() *comments.Domain {
	return &comments.Domain{
	User_id:	req.User_id,
	Thread_id:	req.Thread_id,
	Comment:	req.Comment,
	}
}

func (req *Comment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
