package main

import (
	"log/slog"
	"os"

	"github.com/bmstu-itstech/itsreg-api/internal/config"
	"github.com/bmstu-itstech/itsreg-api/internal/server"
	"github.com/bmstu-itstech/itsreg-api/internal/server/router"
	"github.com/bmstu-itstech/itsreg-api/pkg/logger/handlers/slogpretty"

	_ "github.com/bmstu-itstech/itsreg-api/docs"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal, envDev:
		log = slog.New(
			slogpretty.PrettyHandlerOptions{
				SlogOpts: &slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			}.NewPrettyHandler(os.Stdout),
		)
	case envProd:
		log = slog.New(
			slogpretty.PrettyHandlerOptions{
				SlogOpts: &slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			}.NewPrettyHandler(os.Stdout),
		)
	}

	return log
}

//	@title			ITS Reg API
//	@version		0.1.2

// @contact.name	Zhikharev Kirill
// @contact.url		https://t.me/zhikhkirill
// @contact.email	zhikh.k@gmail.com
// @BasePath  		/api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	app, err := server.New(log, cfg)
	if err != nil {
		panic(err)
	}

	router.Configure(app)

	err = app.Start()
	if err != nil {
		panic(err)
	}
}
