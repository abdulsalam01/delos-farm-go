package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/abdulsalam/delos/config/app"
	_dbHandler "github.com/abdulsalam/delos/config/db"
	"github.com/abdulsalam/delos/helper"
	_middlewareHandler "github.com/abdulsalam/delos/middleware"

	_farmHandler "github.com/abdulsalam/delos/internal/handler/farm/http"
	_farmRepo "github.com/abdulsalam/delos/internal/repository/farm"
	_farmUc "github.com/abdulsalam/delos/internal/usecase/farm"

	_pondHandler "github.com/abdulsalam/delos/internal/handler/pond/http"
	_pondRepo "github.com/abdulsalam/delos/internal/repository/pond"
	_pondUc "github.com/abdulsalam/delos/internal/usecase/pond"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	// Init validator.
	helper.NewValidator()
	// Init middleware statistics.
	_middlewareHandler.InitStatistics()

	// Load appConfig.
	config, err := app.LoadAppConfig()
	if err != nil {
		log.Fatalf("failed init config %v", err)
		return
	}

	// Init database.
	log.Println("Load database")
	db, err := sql.Open("sqlite3", config.Database.Path)
	if err != nil {
		log.Fatalf("failed init database %v", err)
		return
	}

	// Init tables.
	if err := _dbHandler.New(db); err != nil {
		log.Fatalf("failed init tables %v", err)
		return
	}

	// Init repo.
	log.Println("Load Repository")
	farmRepo := _farmRepo.New(db)
	pondRepo := _pondRepo.New(db)

	// Init usecase.
	log.Println("Load usecase")
	farmUsecase := _farmUc.New(farmRepo, pondRepo)
	pondUsecase := _pondUc.New(farmRepo, pondRepo)

	// Init handler.
	log.Println("Load handler")
	farmHandler := _farmHandler.New(farmUsecase)
	pondHandler := _pondHandler.New(farmUsecase, pondUsecase)

	// Routes.
	log.Println("Setup routes")
	r := setupRoutes(chi.NewRouter(), farmHandler, pondHandler)

	// Run apps.
	log.Printf("App start on port :%s", config.App.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), r)
}

func setupRoutes(
	r *chi.Mux,
	farmHandler *_farmHandler.Handler,
	pondHandler *_pondHandler.Handler,
) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(_middlewareHandler.JSONMiddleware)
	r.Use(_middlewareHandler.TrackStatistics)

	// Farm.
	r.Route("/farm", func(r chi.Router) {
		r.Post("/", farmHandler.Create)           // POST /farms
		r.Post("/upsert", farmHandler.Upsert)     // POST /farms/upsert
		r.Delete("/{id}", farmHandler.DeleteByID) // Delete /farms/:id
		r.Put("/{id}", farmHandler.Update)        // PUT /farms/:id
		r.Get("/", farmHandler.GetAll)            // GET /farms
		r.Get("/{id}", farmHandler.GetByID)       // GET /farms/:id
	})

	// Pond.
	r.Route("/pond", func(r chi.Router) {
		r.Post("/", pondHandler.Create)           // POST /ponds
		r.Post("/upsert", pondHandler.Upsert)     // POST /ponds/upsert
		r.Delete("/{id}", pondHandler.DeleteByID) // Delete /ponds/:id
		r.Put("/{id}", pondHandler.Update)        // PUT /ponds/:id
		r.Get("/", pondHandler.GetAll)            // GET /ponds
		r.Get("/{id}", pondHandler.GetByID)       // GET /ponds/:id
	})

	// Statistic.
	r.Get("/statistic", _middlewareHandler.GetStatistics)

	return r
}
