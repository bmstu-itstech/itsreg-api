package handler

import (
	"errors"
	"github.com/bmstu-itstech/itsreg-api/internal/infrastructure/tg"
	"github.com/bmstu-itstech/itsreg-api/internal/server"
	"github.com/bmstu-itstech/itsreg-api/internal/server/request"
	"github.com/bmstu-itstech/itsreg-api/internal/server/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BotsHandler struct {
	s *server.Server
}

func NewBotsHandler(
	s *server.Server,
) *BotsHandler {
	return &BotsHandler{
		s: s,
	}
}

// Create godoc
//
//	@Summary		Create a bot
//	@ID				bot-create
//	@Accept			json
//	@Produce		json
//	@Param			params	body		request.Create	true	"Bot data"
//	@Success		200		{object}	response.Create
//	@Failure		400		{object}	response.Error
//	@Failure		401		{object}	response.Error
//	@Failure		403		{object}	response.Error
//	@Router			/bots/create [post]
//
// @Security BearerAuth
func (h *BotsHandler) Create(c echo.Context) error {
	req := new(request.Create)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
	}

	// TODO: create grpc itsreg.BotsService.Create

	return c.JSON(http.StatusNotImplemented, response.Error{Message: "Not implemented"})
}

// Bot godoc
//
//	@Summary		Get bot model
//	@ID				bot-get
//	@Accept			json
//	@Produce		json
//	@Param          id 		path 		int 			true 	"id"
//	@Success		200		{object}	response.Bot
//	@Failure		400		{object}	response.Error
//	@Failure		401		{object}	response.Error
//	@Failure		403		{object}	response.Error
//	@Failure		404		{object}	response.Error
//	@Router			/bots/{id} [get]
//
// @Security BearerAuth
func (h *BotsHandler) Bot(c echo.Context) error {
	_, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// TODO: create grpc itsreg.BotsService.Bot

	return c.JSON(http.StatusNotImplemented, response.Error{Message: "Not implemented"})
}

// Start godoc
//
//	@Summary		Start telegram bot
//	@ID				bot-start
//	@Accept			json
//	@Produce		json
//	@Param          id 		path 		int 			true 	"id"
//	@Success		204		{object}	response.Empty
//	@Failure		400		{object}	response.Error
//	@Failure		401		{object}	response.Error
//	@Failure		403		{object}	response.Error
//	@Failure		404		{object}	response.Error
//	@Router			/bots/{id}/start [post]
//
// @Security BearerAuth
func (h *BotsHandler) Start(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.s.Tg.Start(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, tg.ErrNotFound) {
			return c.JSON(http.StatusNotFound, response.Error{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Empty{})
}

// Stop godoc
//
//	@Summary		Stop telegram bot
//	@ID				bot-stop
//	@Accept			json
//	@Produce		json
//	@Param          id 		path 		int 			true 	"id"
//	@Success		204		{object}	response.Empty
//	@Failure		400		{object}	response.Error
//	@Failure		401		{object}	response.Error
//	@Failure		403		{object}	response.Error
//	@Failure		404		{object}	response.Error
//	@Router			/bots/{id}/stop [post]
//
// @Security BearerAuth
func (h *BotsHandler) Stop(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = h.s.Tg.Stop(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, tg.ErrNotFound) {
			return c.JSON(http.StatusNotFound, response.Error{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, response.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Empty{})
}
