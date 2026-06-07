package app

import "example/cats/internal/handlers"

func (a *App) loadRoutes() {
	// health handler
	healthHandler := handlers.HealthHandler{
		DB:     a.DB,
		Logger: a.Logger,
	}

	a.Router.HandleFunc("GET /health", healthHandler.Health())

}
