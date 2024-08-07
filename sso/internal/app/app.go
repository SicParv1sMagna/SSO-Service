package app

import (
	"log/slog"
	"time"

	app "github.com/SicParv1sMagna/sso-service/sso/internal/app/grpc"
	"github.com/SicParv1sMagna/sso-service/sso/internal/services/auth"
	"github.com/SicParv1sMagna/sso-service/sso/internal/storage/sqlite"
)

type App struct {
	GRPCServer *app.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	// This seems strange, but it gives flexibility
	// Jump into auth package implementation
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := app.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
