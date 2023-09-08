package main

import (
	"log"
	"net/http"

	"bookstore.com/api"
	"bookstore.com/config"
	"bookstore.com/domain"
	"bookstore.com/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := repository.NewAuthorMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := domain.NewAuthorService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/authors/{code}", handler.Get)
	r.Post("/authors", handler.Post)
	r.Delete("/authors/{code}", handler.Delete)
	r.Get("/authors", handler.GetAll)
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))
}
