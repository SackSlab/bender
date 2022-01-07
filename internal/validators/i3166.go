package validators

import (
	"fmt"
	"net/http"

	"github.com/sackslab/bender/internal/i3166"
	"github.com/sackslab/bender/internal/middlewares/apperror"
)

type I3166 i3166.ISO3166

func (i *I3166) UnmarshalJSON(data []byte) error {
	code := string([]rune(string(data)))
	code = code[1 : len(code)-1]
	i3166, found := i3166.ByCode(code)

	if !found {
		return &apperror.AppError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("invalid country code or type %s,  expected ISO3166 values", code),
		}
	}

	i.Code, i.Name = i3166.Code, i3166.Name

	return nil
}
