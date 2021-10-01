package deliveries

import (
	"net/http"

	"github.com/NunChatSpace/identity-service/internal/response_wrapper"
)

func InternalError(model interface{}, err error) (response_wrapper.Model, error) {
	return response_wrapper.Model{
		StatusCode: http.StatusInternalServerError,
		Data:       model,
		Message:    "",
	}, err
}

func Forbidden(model interface{}, err error) (response_wrapper.Model, error) {
	return response_wrapper.Model{
		StatusCode: http.StatusInternalServerError,
		Data:       model,
		Message:    "",
	}, err
}
