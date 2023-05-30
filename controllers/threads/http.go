package threads

import (
	"Capstone/controllers"
	"Capstone/controllers/threads/request"
	"Capstone/controllers/threads/response"
	"Capstone/models/threads"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadController struct {
	threadUsecase threads.Usecase
}

func NewThreadlController(threadUC threads.Usecase) *ThreadController {
	return &ThreadController{
		threadUsecase: threadUC,
	}
}

func (cc *ThreadController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	threadsData, err := cc.threadUsecase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	threads := []response.Thread{}

	for _, thread := range threadsData {
		threads = append(threads, response.FromDomain(thread))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all products", threads)
}

func (cc *ThreadController) GetByID(c echo.Context) error {
	var threadID string = c.Param("id")
	ctx := c.Request().Context()

	thread, err := cc.threadUsecase.GetByID(ctx, threadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "product not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product found", response.FromDomain(thread))
}

func (cc *ThreadController) GetByName(c echo.Context) error {
	var threadName string = c.Param("name")
	ctx := c.Request().Context()

	thread, err := cc.threadUsecase.GetByName(ctx, threadName)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "Product not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product found", response.FromDomain(thread))
}

func (cc *ThreadController) Create(c echo.Context) error {
	input := request.Thread{}
	ctx := c.Request().Context()

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	thread, err := cc.threadUsecase.Create(ctx, input.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a product", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "product's created", response.FromDomain(thread))
}

func (cc *ThreadController) Update(c echo.Context) error {
	var threadID string = c.Param("id")
	ctx := c.Request().Context()

	input := request.Thread{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	err := input.Validate()

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	thread, err := cc.threadUsecase.Update(ctx, input.ToDomain(), threadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update product failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product updated", response.FromDomain(thread))
}

func (cc *ThreadController) Delete(c echo.Context) error {
	var threadID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.threadUsecase.Delete(ctx, threadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete product failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product deleted", "")
}

func (cc *ThreadController) Restore(c echo.Context) error {
	var threadID string = c.Param("id")
	ctx := c.Request().Context()

	thread, err := cc.threadUsecase.Restore(ctx, threadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "product not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product restored", response.FromDomain(thread))
}

func (cc *ThreadController) ForceDelete(c echo.Context) error {
	var threadID string = c.Param("id")
	ctx := c.Request().Context()

	err := cc.threadUsecase.ForceDelete(ctx, threadID)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete product failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "product deleted permanently", "")
}
