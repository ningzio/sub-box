package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ningzio/sub-box/backend/internal/config"
	"github.com/ningzio/sub-box/backend/pkg/schema"
	"github.com/sagernet/serenity/option"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"http://localhost:5173"}, // Vite 默认端口
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/schema", s.handleGetSchema())
	})

	return r
}

func (s *Server) Foo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}
}

func (s *Server) handleGetSchema() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		generator := schema.NewGenerator(schema.SchemaOptions{
			RequireAll:       false,
			IgnoreZeroValues: true,
			IncludeExamples:  true,
			Registry:         config.NewRegistry(),
		})

		schema, err := generator.Generate(option.Template{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(schema)
	}
}
