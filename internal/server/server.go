package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fekuna/go-post-article/config"
	"github.com/fekuna/go-post-article/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
	db     *sqlx.DB
}

// NewServer new Server constructor
func NewServer(cfg *config.Config, logger logger.Logger, db *sqlx.DB) *Server {
	return &Server{
		echo:   echo.New(),
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("Error starting server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Println("Server exited properly")
	return s.echo.Server.Shutdown(ctx)
}
