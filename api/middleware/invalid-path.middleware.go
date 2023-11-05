package middleware

import (
	"github.com/labstack/echo/v4"
	error_response "github.com/nguyendhst/fx-template/domain/response/error"
)

func InvalidPathResponseFormatMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Call the next handler
		err := next(c)
		if err != nil {
			return error_response.DataNotFoundError(
				c,
				error_response.PATH_NOT_FOUND_MESSAGE,
				error_response.PATH_NOT_FOUND_CODE,
			)
		}

		return nil
	}
}
