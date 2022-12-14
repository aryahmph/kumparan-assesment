package http

import (
	"errors"
	errorCommon "github.com/aryahmph/kumparan-assesment/common/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func MiddlewareErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs[0]
			if !err.IsType(gin.ErrorTypePrivate) {
				var ves validator.ValidationErrors
				if errors.As(err, &ves) {
					keys := make(map[string]string)
					for _, ve := range ves {
						keys[ve.Field()] = ve.Tag()
					}
					c.JSON(c.Writer.Status(), Error{
						Code:    c.Writer.Status(),
						Message: err.Error(),
						Errors:  keys,
					})
					return
				}
			}

			switch err := err.Err.(type) {
			case *errorCommon.ClientError:
				c.JSON(err.StatusCode, Error{
					Code:    err.StatusCode,
					Message: err.Message,
				})
			default:
				c.JSON(http.StatusInternalServerError, Error{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				})
			}
		}
	}
}
