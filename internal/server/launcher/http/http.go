package http

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"net/http"
	"notify-service/internal/server/launcher"
	"time"
)

type server struct {
	srv *http.Server
}

func New(handler http.Handler) launcher.Server {
	var httpServer = &http.Server{
		Addr:           "localhost" + ":" + viper.GetString("http_port"),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	srv := &server{
		srv: httpServer,
	}

	return srv
}

func (s *server) Serve(ctx context.Context) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		if err := s.srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():

	case err := <-errCh:
		return err
	}

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	// TODO: modify graceful shutdown with error return

	return nil
}
