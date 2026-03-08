package server

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rdl/core/internal/platform/config"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	cfg    config.ServerConfig
	log    *slog.Logger
	http   *http.Server
	mongo  *mongo.Client
	pg     *sql.DB
	redis  *redis.Client
	nats   *nats.Conn
}

func New(
	cfg config.ServerConfig,
	log *slog.Logger,
	mongo *mongo.Client,
	pg *sql.DB,
	rdb *redis.Client,
	nc *nats.Conn,
) *Server {
	s := &Server{
		cfg:   cfg,
		log:   log,
		mongo: mongo,
		pg:    pg,
		redis: rdb,
		nats:  nc,
	}

	s.http = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:      s.routes(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return s
}

func (s *Server) Start() error {
	s.log.Info("server starting", "addr", s.http.Addr)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", s.handleHealth)

	return mux
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
