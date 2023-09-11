package register

import (
	"fmt"

	"github.com/labstack/echo/v4"
	resperr "github.com/nguyendhst/lagile/domain/response/error"
	domain "github.com/nguyendhst/lagile/domain/user"
	"github.com/nguyendhst/lagile/module/config"
)

type (
	RegisterController struct {
		RegisterUsecase domain.RegisterUsecase
		Env             *config.Env
	}
)

func New(registerUsecase domain.RegisterUsecase, env *config.Env) *RegisterController {
	return &RegisterController{
		RegisterUsecase: registerUsecase,
		Env:             env,
	}
}

func (rc *RegisterController) RegisterUser(c echo.Context) error {
	// Get request body
	req := new(domain.UserRegisterRequest)
	if err := c.Bind(req); err != nil {
		return resperr.BadRequestError(c, resperr.REGISTER_BAD_REQUEST_MESSAGE, resperr.REGISTER_ERROR_CODE)
	}

	// Call usecase
	user, err := rc.RegisterUsecase.RegisterUser(c.Request().Context(), &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return resperr.InternalServerError(c, resperr.INTERNAL_SERVER_ERROR_MESSAGE, resperr.INTERNAL_SERVER_ERROR_CODE)
	}

	res := rc.RegisterUsecase.NewResponse(&domain.UserRegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})

	fmt.Println(res)

	return c.JSON(200, res)
}
