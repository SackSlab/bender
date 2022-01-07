package apperror

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	err     error  `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *AppError) Error() string {
	return fmt.Sprintf("AppError{msg:%s  code:%d  originalMsg: %s}", err.Message, err.Code, err.err.Error())
}

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err

			if aerr, ok := err.(*AppError); ok {
				c.AbortWithStatusJSON(aerr.Code, aerr)
				return
			}

			if err == io.EOF {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, &AppError{
					Code:    http.StatusUnsupportedMediaType,
					Message: "invalid mime type",
				})
				return
			}

			if ute, ok := err.(*json.UnmarshalTypeError); ok {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, &AppError{
					err:     err,
					Code:    http.StatusBadRequest,
					Message: fmt.Sprintf("Field: '%s' required type is '%s'", ute.Field, ute.Value),
				})
				return
			}

			// TODO: format message of validation erros
			if validationErrors, ok := err.(validator.ValidationErrors); ok && len(validationErrors) > 0 {
				parsedErrs := make([]*AppError, 0)
				for _, v := range validationErrors {
					parsedErrs = append(parsedErrs, &AppError{
						err:     v,
						Code:    http.StatusBadRequest,
						Message: v.Error(),
					})
				}

				c.AbortWithStatusJSON(http.StatusInternalServerError, parsedErrs)
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, &AppError{
				err:     err,
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
			})
		}
	}
}
