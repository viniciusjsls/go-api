package app

import (
	"net/http"
	"time"

	"github.com/viniciusjsls/go-api/internal/config"
	"go.uber.org/zap"
)

func NewServer(cfg *config.Config, logger *zap.Logger) *http.Server {
	router := InitRouter(cfg, logger)

	return &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}
