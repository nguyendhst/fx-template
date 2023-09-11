package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nguyendhst/clean-architecture-skeleton/domain/auth"
	resperr "github.com/nguyendhst/clean-architecture-skeleton/domain/response/error"
	"github.com/nguyendhst/clean-architecture-skeleton/module/config"
	"github.com/nguyendhst/clean-architecture-skeleton/shared/util"
)

type LoginController struct {
	LoginUsecase auth.LoginUsecase
	Env          *config.Env
}

func New(env *config.Env, loginUsecase auth.LoginUsecase) *LoginController {
	return &LoginController{
		LoginUsecase: loginUsecase,
		Env:          env,
	}
}

func (controller *LoginController) Login(c echo.Context) error {
	// Get request body
	req := new(auth.LoginRequest)
	if err := c.Bind(req); err != nil {
		return resperr.BadRequestError(c, resperr.LOGIN_BAD_REQUEST_MESSAGE, resperr.LOGIN_ERROR_CODE)
	}

	// Call usecase
	user, err := controller.LoginUsecase.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return resperr.DataNotFoundError(c, resperr.DATA_NOT_FOUND_MESSAGE, resperr.DATA_NOT_FOUND_CODE)
	}

	// Check password hash
	if !util.CheckPasswordHash(req.Password, user.Password) {
		return resperr.UnauthorizedError(c, resperr.INVALID_CREDENTIALS_UNAUTHORIZED_MESSAGE, resperr.INVALID_CREDENTIALS_UNAUTHORIZED_CODE)
	}

	// Create access token
	accessToken, err := controller.LoginUsecase.CreateAccessToken(&user, controller.Env.AccessTokenSecret, controller.Env.AccessTokenExpiryHour)
	if err != nil {
		return resperr.InternalServerError(c, resperr.INTERNAL_SERVER_ERROR_MESSAGE, resperr.INTERNAL_SERVER_ERROR_CODE)
	}

	// Create refresh token
	refreshToken, err := controller.LoginUsecase.CreateRefreshToken(&user, controller.Env.RefreshTokenSecret, controller.Env.RefreshTokenExpiryHour)
	if err != nil {
		return resperr.InternalServerError(c, resperr.INTERNAL_SERVER_ERROR_MESSAGE, resperr.INTERNAL_SERVER_ERROR_CODE)
	}

	// Create response
	response := auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Return response
	return c.JSON(http.StatusOK, response)
}
