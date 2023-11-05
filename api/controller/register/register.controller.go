package register

import (
	"github.com/labstack/echo/v4"
	resperr "github.com/nguyendhst/fx-template/domain/response/error"
	domain "github.com/nguyendhst/fx-template/domain/user"
	"github.com/nguyendhst/fx-template/module/config"
)

type (
	RegisterController struct {
		RegisterUsecase domain.RegisterUsecase
		Configs         *config.Config
	}
)

func New(registerUsecase domain.RegisterUsecase, config *config.Config) *RegisterController {
	return &RegisterController{
		RegisterUsecase: registerUsecase,
		Configs:         config,
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

	res := domain.UserRegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(200, res)
}
