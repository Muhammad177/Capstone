package likes

import (
	"Capstone/controllers"
	"Capstone/controllers/likes/request"
	"Capstone/controllers/likes/response"
	"Capstone/models/likes"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeUsecase likes.Usecase
}

func NewLikelController(likeUC likes.Usecase) *LikeController {
	return &LikeController{
		likeUsecase: likeUC,
	}
}

func (cc *LikeController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	likesData, err := cc.likeUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	likes := []response.Like{}

	for _, like := range likesData {
		likes = append(likes, response.FromDomain(like))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "see all likes", likes)
}

func (cc *LikeController) GetByID(c echo.Context) error {
	var likeID string = c.Param("id")
	ctx := c.Request().Context()

	like, err := cc.likeUsecase.GetByID(ctx, likeID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like found", response.FromDomain(like))
}

func (cc *LikeController) GetByName(c echo.Context) error {
	var likeName string = c.Param("name")
	ctx := c.Request().Context()

	like, err := cc.likeUsecase.GetByName(ctx, likeName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like found", response.FromDomain(like))
}

func (cc *LikeController) Create(c echo.Context) error {
	input := request.Like{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	like, err := cc.likeUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a like", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "like's created", response.FromDomain(like))
}

func (cc *LikeController) Update(c echo.Context) error {
	var likeID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Like{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	like, err := cc.likeUsecase.Update(ctx, input.ToDomain(), likeID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update like failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like updated", response.FromDomain(like))
}

func (cc *LikeController) Delete(c echo.Context) error {
	var likeID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.likeUsecase.Delete(ctx, likeID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete like failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like deleted", "")
}

func (cc *LikeController) Restore(c echo.Context) error {
	var likeID string = c.Param("id")
	ctx := c.Request().Context()

	like, err := cc.likeUsecase.Restore(ctx, likeID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "like not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like restored", response.FromDomain(like))
}

func (cc *LikeController) ForceDelete(c echo.Context) error {
	var likeID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.likeUsecase.ForceDelete(ctx, likeID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete like failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "like deleted permanently", "")
}
