package main

import (
	"github.com/AidarBabanov/billups-interview/config"
	"github.com/AidarBabanov/billups-interview/game"
	"github.com/AidarBabanov/billups-interview/logging"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kelseyhightower/envconfig"
	logs "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// load app configuration from environment variables
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		logs.Fatal(err)
		os.Exit(-1)
	}
	logging.Init(cfg.LogLevel)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", http.RedirectHandler("/doc/", http.StatusMovedPermanently).ServeHTTP)
	router.Get("/doc/*", http.StripPrefix("/doc/", http.FileServer(http.Dir("./static"))).ServeHTTP)

	gameService := game.NewGame()
	router.Mount("/", game.NewResource(gameService))

	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.ServerPort),
		Handler:      router,
		WriteTimeout: cfg.ServerWriteTimeout,
		ReadTimeout:  cfg.ServerReadTimeout,
	}

	logs.Info("Serving the web application...")
	logs.Fatal(httpServer.ListenAndServe())
}
