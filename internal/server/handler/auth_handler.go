package handler

import (
	"github.com/bmstu-itstech/itsreg-api/internal/server"
	"github.com/bmstu-itstech/itsreg-api/internal/server/request"
	"github.com/bmstu-itstech/itsreg-api/internal/server/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler struct {
	s *server.Server
}

func NewAuthHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}

// Register godoc
//
//	@Summary		Register a user
//	@ID				auth-register
//	@Accept			json
//	@Produce		json
//	@Param			params	body		request.Register	true	"User's credentials"
//	@Success		200		{object}	response.Authorized
//	@Failure		400		{object}	response.Error
//	@Failure		409		{object}	response.Error
//	@Router			/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	req := new(request.Register)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: create grpc itsreg.AuthService

	return c.JSON(http.StatusNotImplemented, response.Error{Message: "Not implemented"})
}

// Login godoc
//
//	@Summary		Authenticate a user
//	@ID				auth-login
//	@Accept			json
//	@Produce		json
//	@Param			params	body		request.Login	true	"User's credentials"
//	@Success		200		{object}	response.Authorized
//	@Failure		400		{object}	response.Error
//	@Failure		401		{object}	response.Error
//	@Router			/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	req := new(request.Login)

	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: create grpc itsreg.AuthService

	return c.JSON(http.StatusNotImplemented, response.Error{Message: "Not implemented"})
}
