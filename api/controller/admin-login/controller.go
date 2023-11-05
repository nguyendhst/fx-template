package admin_login

import (
	"net/http"

	"github.com/nguyendhst/fx-template/domain/auth"
	"github.com/nguyendhst/fx-template/module/config"

	"github.com/labstack/echo/v4"
	response_error "github.com/nguyendhst/fx-template/domain/response/error"

	"github.com/nguyendhst/fx-template/shared/util"
)

type (
	AdminLoginController struct {
		LoginUsecase auth.BasicLoginUsecase
		Configs      *config.Config
	}
)

func New(configs *config.Config, loginUsecase auth.BasicLoginUsecase) *AdminLoginController {
	return &AdminLoginController{
		LoginUsecase: loginUsecase,
		Configs:      configs,
	}
}

func (controller *AdminLoginController) Login(c echo.Context) error {
	// Get request body
	req := new(auth.BasicLoginRequest)
	if err := c.Bind(req); err != nil {
		return response_error.BadRequestError(
			c,
			response_error.LOGIN_BAD_REQUEST_MESSAGE,
			response_error.LOGIN_ERROR_CODE,
		)
	}

	// Call usecase
	user, err := controller.LoginUsecase.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return response_error.DataNotFoundError(
			c,
			response_error.DATA_NOT_FOUND_MESSAGE,
			response_error.DATA_NOT_FOUND_CODE,
		)
	}

	// Check password hash
	if !util.CheckPasswordHash(req.Password, user.Password) {
		return response_error.UnauthorizedError(
			c,
			response_error.INVALID_CREDENTIALS_UNAUTHORIZED_MESSAGE,
			response_error.INVALID_CREDENTIALS_UNAUTHORIZED_CODE,
		)
	}

	// Create access token
	accessToken, err := controller.LoginUsecase.CreateAccessToken(
		&user,
		controller.Configs.Env.Secret.JwtSecret.Access.Key,
		controller.Configs.Env.Secret.JwtSecret.Access.Expiration)
	if err != nil {
		return response_error.InternalServerError(
			c,
			response_error.INTERNAL_SERVER_ERROR_MESSAGE,
			response_error.INTERNAL_SERVER_ERROR_CODE,
		)
	}

	// Create refresh token
	refreshToken, err := controller.LoginUsecase.CreateRefreshToken(
		&user,
		controller.Configs.Env.Secret.JwtSecret.Refresh.Key,
		controller.Configs.Env.Secret.JwtSecret.Refresh.Expiration,
	)
	if err != nil {
		return response_error.InternalServerError(
			c,
			response_error.INTERNAL_SERVER_ERROR_MESSAGE,
			response_error.INTERNAL_SERVER_ERROR_CODE,
		)
	}

	// Create response
	response := auth.BasicLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Return response
	return c.JSON(http.StatusOK, response)
}
