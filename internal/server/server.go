package server

import (
	"github.com/bmstu-itstech/itsreg-api/internal/config"
	"github.com/bmstu-itstech/itsreg-api/internal/infrastructure/bots"
	"github.com/bmstu-itstech/itsreg-api/internal/infrastructure/tg"
	"github.com/bmstu-itstech/itsreg-api/pkg/endpoint"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Server struct {
	Echo   *echo.Echo
	Log    *slog.Logger
	Config *config.Config

	Bots *bots.Service
	Tg   *tg.Service
}

func New(
	log *slog.Logger,
	cfg *config.Config,
) (*Server, error) {
	botsService, err := bots.NewService(log, cfg.BotsConfig.Host, cfg.BotsConfig.Port)
	if err != nil {
		return nil, err
	}

	tgService, err := tg.NewService(log, cfg.TgConfig.Host, cfg.TgConfig.Port)
	if err != nil {
		return nil, err
	}

	return &Server{
		Echo:   echo.New(),
		Config: cfg,
		Bots:   botsService,
		Tg:     tgService,
	}, nil
}

func (s *Server) Start() error {
	addr := endpoint.Addr(s.Config.Server.Host, s.Config.Server.Port)
	return s.Echo.Start(addr)
}
