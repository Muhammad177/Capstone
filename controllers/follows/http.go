package follows

import (
	"Capstone/controllers"
	"Capstone/controllers/follows/request"
	"Capstone/controllers/follows/response"
	"Capstone/models/follows"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FollowController struct {
	followUsecase follows.Usecase
}

func NewFollowlController(followUC follows.Usecase) *FollowController {
	return &FollowController{
		followUsecase: followUC,
	}
}

func (cc *FollowController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	followsData, err := cc.followUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	follows := []response.Follow{}

	for _, follow := range followsData {
		follows = append(follows, response.FromDomain(follow))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get all follows", follows)
}

func (cc *FollowController) GetByID(c echo.Context) error {
	var followID string = c.Param("id")
	ctx := c.Request().Context()

	follow, err := cc.followUsecase.GetByID(ctx, followID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "follow not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow found", response.FromDomain(follow))
}

func (cc *FollowController) GetByName(c echo.Context) error {
	var followName string = c.Param("name")
	ctx := c.Request().Context()

	follow, err := cc.followUsecase.GetByName(ctx, followName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "follow not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow found", response.FromDomain(follow))
}

func (cc *FollowController) Create(c echo.Context) error {
	input := request.Follow{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	follow, err := cc.followUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a follow", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "follow's created", response.FromDomain(follow))
}

func (cc *FollowController) Update(c echo.Context) error {
	var followID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Follow{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	follow, err := cc.followUsecase.Update(ctx, input.ToDomain(), followID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update follow failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow updated", response.FromDomain(follow))
}

func (cc *FollowController) Delete(c echo.Context) error {
	var followID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.followUsecase.Delete(ctx, followID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete follow failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow deleted", "")
}

func (cc *FollowController) Restore(c echo.Context) error {
	var followID string = c.Param("id")
	ctx := c.Request().Context()

	follow, err := cc.followUsecase.Restore(ctx, followID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "follow not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow restored", response.FromDomain(follow))
}

func (cc *FollowController) ForceDelete(c echo.Context) error {
	var followID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.followUsecase.ForceDelete(ctx, followID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete follow failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "follow deleted permanently", "")
}
