package handler

import (
	"net/http"

	"github.com/richhh7g/mm-api-nfe/internal/api/dto"
	"github.com/richhh7g/mm-api-nfe/internal/api/response"
)

// Get v1 hello world
// @Summary Get hello world message
// @Tags v1, Hello World
// @Produce json
// @Success 200 {object} GetHelloWorldResponse
// @Router /api/v1/hello-world [GET]
func GetV1HelloWorld(w http.ResponseWriter, r *http.Request) {
	helloResponse := dto.GetHelloWorldResponse{
		Message: "Hello World!",
	}

	response.NewSuccess(helloResponse, http.StatusOK).Send(w)
}
