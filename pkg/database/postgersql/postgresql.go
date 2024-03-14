package postgersql

import (
	"database/sql"
	"fmt"

	"github.com/abylaymoldabek/finSystem/internal/config"
	"github.com/abylaymoldabek/finSystem/pkg/logger"
	_ "github.com/lib/pq"
)

func NewClient(cfg *config.Config) (*sql.DB, error) {
	logger.Info("Database connected!")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBConfig.Host, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.DBName)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return db, nil
}
