package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/krossroad/imperialfleet/internal/imperialfleet"
	"github.com/krossroad/imperialfleet/internal/logger"
	"github.com/krossroad/imperialfleet/internal/sql"
)

type (
	Service struct {
		db         *gorm.DB
		log        *logger.Entry
		httpServer *http.Server
	}
	Options struct {
		DBUrl    string
		HTTPAddr string
	}
)

func (s *Service) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

func NewService(ctx context.Context, log *logger.Entry, opt Options) (*Service, error) {
	db, err := gorm.Open(mysql.Open(opt.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	svc := &Service{
		log: log,
		db:  db,
	}

	sqlPersist := sql.New(db)
	r := mux.NewRouter()
	mySvc := imperialfleet.New(log, sqlPersist)

	// Register routes
	registerRoute(r, mySvc)

	httpServer := &http.Server{
		Addr:    opt.HTTPAddr,
		Handler: r,
	}
	svc.httpServer = httpServer

	go func() {
		svc.log.With("address", opt.HTTPAddr).Info("starting http-server")
		if err := svc.httpServer.ListenAndServe(); err != nil {
			svc.log.Fatal("unable serve http", "error", err)
		}
	}()

	return svc, nil
}

func registerRoute(r *mux.Router, svc *imperialfleet.Service) {
	r.HandleFunc("/", svc.Create).Methods(http.MethodPost)
	r.HandleFunc("/", svc.List).Methods(http.MethodGet)
	r.HandleFunc("/{id}", svc.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/{id}", svc.Update).Methods(http.MethodPut)
	r.HandleFunc("/{id}", svc.Show).Methods(http.MethodGet)
}
