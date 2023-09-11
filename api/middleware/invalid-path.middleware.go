package middleware

import (
	"github.com/labstack/echo/v4"
	resperr "github.com/nguyendhst/clean-architecture-skeleton/domain/response/error"
)

func InvalidPathResponseFormatMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Call the next handler
		err := next(c)
		if err != nil {
			return resperr.DataNotFoundError(
				c,
				resperr.PATH_NOT_FOUND_MESSAGE,
				resperr.PATH_NOT_FOUND_CODE,
			)
		}

		return nil
	}
}
