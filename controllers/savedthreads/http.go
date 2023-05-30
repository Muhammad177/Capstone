package savedthreads

import (
	"Capstone/controllers"
	"Capstone/controllers/savedthreads/request"
	"Capstone/controllers/savedthreads/response"
	"Capstone/models/savedthreads"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SavedthreadController struct {
	savedthreadUsecase savedthreads.Usecase
}

func NewSavedthreadlController(savedthreadUC savedthreads.Usecase) *SavedthreadController {
	return &SavedthreadController{
		savedthreadUsecase: savedthreadUC,
	}
}

func (cc *SavedthreadController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	savedthreadsData, err := cc.savedthreadUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	savedthreads := []response.Savedthread{}

	for _, savedthread := range savedthreadsData {
		savedthreads = append(savedthreads, response.FromDomain(savedthread))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all threads saved", savedthreads)
}

func (cc *SavedthreadController) GetByID(c echo.Context) error {
	var savedthreadID string = c.Param("id")
	ctx := c.Request().Context()

	savedthread, err := cc.savedthreadUsecase.GetByID(ctx, savedthreadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "threads not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "threads found", response.FromDomain(savedthread))
}

func (cc *SavedthreadController) GetByName(c echo.Context) error {
	var savedthreadName string = c.Param("name")
	ctx := c.Request().Context()

	savedthread, err := cc.savedthreadUsecase.GetByName(ctx, savedthreadName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "Threads not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Threads found", response.FromDomain(savedthread))
}

func (cc *SavedthreadController) Create(c echo.Context) error {
	input := request.Savedthread{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	savedthread, err := cc.savedthreadUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to save a thread", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "thread's savedd", response.FromDomain(savedthread))
}

func (cc *SavedthreadController) Update(c echo.Context) error {
	var savedthreadID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Savedthread{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	savedthread, err := cc.savedthreadUsecase.Update(ctx, input.ToDomain(), savedthreadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update Threads failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Threads updated", response.FromDomain(savedthread))
}

func (cc *SavedthreadController) Delete(c echo.Context) error {
	var savedthreadID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.savedthreadUsecase.Delete(ctx, savedthreadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete Thread failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Thread deleted", "")
}

func (cc *SavedthreadController) Restore(c echo.Context) error {
	var savedthreadID string = c.Param("id")
	ctx := c.Request().Context()

	savedthread, err := cc.savedthreadUsecase.Restore(ctx, savedthreadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "Thread not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Thread restored", response.FromDomain(savedthread))
}

func (cc *SavedthreadController) ForceDelete(c echo.Context) error {
	var savedthreadID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.savedthreadUsecase.ForceDelete(ctx, savedthreadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete Thread failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Thread deleted permanently", "")
}
