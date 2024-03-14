package app

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/abylaymoldabek/finSystem/internal/config"
	delivery "github.com/abylaymoldabek/finSystem/internal/delivery/http"
	"github.com/abylaymoldabek/finSystem/internal/repository"
	"github.com/abylaymoldabek/finSystem/internal/server"
	"github.com/abylaymoldabek/finSystem/internal/service"
	"github.com/abylaymoldabek/finSystem/pkg/database/postgersql"
	"github.com/abylaymoldabek/finSystem/pkg/logger"
)

func Run(configPath string) {
	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	db, err := postgersql.NewClient(&config)
	if err != nil {
		logger.Error(err)
		return
	}
	defer db.Close()
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handler := delivery.NewHandler(services)
	srv := server.NewServer(&config, handler.Init(&config))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Info("Server started!")
	time.Sleep(time.Second * 1)

	select {}

	// const timeout = 5 * time.Second

	// ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	// defer shutdown()
	// if err := srv.Stop(ctx); err != nil {
	// 	logger.Errorf("failed to stop server: %v", err)
	// }

}
