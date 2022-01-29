package http

import (
	"fmt"
	"net/http"

	"github.com/WelcomerTeam/Discord/structs"
	jsoniter "github.com/json-iterator/go"
)

// RestError contains the error structure that is returned by discord.
type RestError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte

	Message *structs.ErrorMessage
}

func NewRestError(req *http.Request, resp *http.Response, body []byte) *RestError {
	var errorMessage structs.ErrorMessage

	_ = jsoniter.Unmarshal(body, errorMessage)

	return &RestError{
		Request:      req,
		Response:     resp,
		ResponseBody: body,
		Message:      &errorMessage,
	}
}

func (r *RestError) Error() string {
	return fmt.Sprintf("%s: %s", r.Response.Status, r.Message.Message)
}
