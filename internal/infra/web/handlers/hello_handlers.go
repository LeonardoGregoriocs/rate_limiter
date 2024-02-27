package handlers

import (
	"net/http"

	responsehandler "github.com/leonardogregoriocs/rate_limiter/internal/pkg/responsehandles"
)

type HelloWebHandlerInterface interface {
	SayHello(w http.ResponseWriter, r *http.Request)
}

type HelloWebHandler struct {
	ResponseHandler responsehandler.WebResponseHandlerInterface
}

func NewHelloWebHandler(rh responsehandler.WebResponseHandlerInterface) *HelloWebHandler {
	return &HelloWebHandler{
		ResponseHandler: rh,
	}
}

func (h *HelloWebHandler) SayHello(w http.ResponseWriter, r *http.Request) {
	h.ResponseHandler.Respond(w, http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}
