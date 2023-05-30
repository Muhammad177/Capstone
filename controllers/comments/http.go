package comments

import (
	"Capstone/controllers"
	"Capstone/controllers/comments/request"
	"Capstone/controllers/comments/response"
	"Capstone/models/comments"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	commentUsecase comments.Usecase
}

func NewCommentlController(commentUC comments.Usecase) *CommentController {
	return &CommentController{
		commentUsecase: commentUC,
	}
}

func (cc *CommentController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	commentsData, err := cc.commentUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	comments := []response.Comment{}

	for _, comment := range commentsData {
		comments = append(comments, response.FromDomain(comment))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get all comments", comments)
}

func (cc *CommentController) GetByID(c echo.Context) error {
	var commentID string = c.Param("id")
	ctx := c.Request().Context()

	comment, err := cc.commentUsecase.GetByID(ctx, commentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "comment found", response.FromDomain(comment))
}

func (cc *CommentController) GetByName(c echo.Context) error {
	var commentName string = c.Param("name")
	ctx := c.Request().Context()

	comment, err := cc.commentUsecase.GetByName(ctx, commentName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "omment found", response.FromDomain(comment))
}

func (cc *CommentController) Create(c echo.Context) error {
	input := request.Comment{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	comment, err := cc.commentUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a comment", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "comment's created", response.FromDomain(comment))
}

func (cc *CommentController) Update(c echo.Context) error {
	var commentID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Comment{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	comment, err := cc.commentUsecase.Update(ctx, input.ToDomain(), commentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update comment failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "comment updated", response.FromDomain(comment))
}

func (cc *CommentController) Delete(c echo.Context) error {
	var commentID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.commentUsecase.Delete(ctx, commentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete comment failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "comment deleted", "")
}

func (cc *CommentController) Restore(c echo.Context) error {
	var commentID string = c.Param("id")
	ctx := c.Request().Context()

	comment, err := cc.commentUsecase.Restore(ctx, commentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "comment not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "comment restored", response.FromDomain(comment))
}

func (cc *CommentController) ForceDelete(c echo.Context) error {
	var commentID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.commentUsecase.ForceDelete(ctx, commentID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete comment failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "comment deleted permanently", "")
}
