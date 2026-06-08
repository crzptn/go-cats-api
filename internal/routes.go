package app

import "example/cats/internal/handlers"

func (a *App) loadRoutes() {
	// health handler
	healthHandler := handlers.HealthHandler{
		DB:     a.DB,
		Logger: a.Logger,
	}

	a.Router.HandleFunc("GET /health", healthHandler.Health())

	// cats handler
	catsHandler := handlers.CatsHandler{
		DB:     a.DB,
		Logger: a.Logger,
	}
	a.Router.HandleFunc("GET /cats", catsHandler.GetAllCats())
	a.Router.HandleFunc("POST /cats", catsHandler.CreateCat())
	a.Router.HandleFunc("GET /cats/{id}", catsHandler.ReadCat())
	a.Router.HandleFunc("PATCH /cats/{id}", catsHandler.UpdateCat())
	a.Router.HandleFunc("DELETE /cats/{id}", catsHandler.DeleteCat())

}
