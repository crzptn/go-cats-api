package app

import (
	"context"
	"example/cats/db"
	"example/cats/internal/middleware"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type App struct {
	DB     *db.Queries
	Addr   string
	Logger *slog.Logger
	Router *http.ServeMux
}

func New(db *db.Queries, Addr string) *App {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &App{
		Router: http.NewServeMux(),
		DB:     db,
		Addr:   Addr,
		Logger: logger,
	}
}

func (a *App) Run(ctx context.Context) error {

	a.loadRoutes()

	router := middleware.CreateMiddlewareStack(
		middleware.Logger(a.Logger, a.Router),
		middleware.NormalizePath,
		middleware.RecoverMiddleware,
	)

	server := http.Server{
		Handler: router,
		Addr:    a.Addr,
	}

	errChan := make(chan error, 1)
	go func() {
		a.Logger.Info("Server starting ", slog.String("url", a.Addr))
		err := server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		a.Logger.Info("Server shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}
	}
	return nil
}
