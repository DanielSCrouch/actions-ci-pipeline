package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Status struct {
	Status string
}

type App struct {
	HostAddr string
	HostPort int
	server   *http.Server
}

func (s *App) Start() {

	addr := fmt.Sprintf("%s:%d", s.HostAddr, s.HostPort)
	s.server = &http.Server{Addr: addr}
	http.HandleFunc("/status", status)

	log.Printf("Starting server: %s", addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed to start, error: %s", err)
	}

}

func (s *App) Stop(ctx context.Context) {

	stop := make(chan struct{})
	go func() {
		if err := s.server.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shutdown server, error: %s", err)
		}

		close(stop)
	}()

	select {
	case <-ctx.Done():
		log.Printf("Shutdown timeout: %v", ctx.Err())
	case <-stop:
		log.Print("Server shutdown.")
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	status := Status{Status: "online"}
	resp, err := json.Marshal(status)
	if err != nil {
		log.Printf("failed to marshal response, error: %s", err)
	}

	fmt.Fprintf(w, string(resp))
}
