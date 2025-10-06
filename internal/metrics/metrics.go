package metrics

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	logger       *slog.Logger
	healthz      bool
	ready        bool
	healthzMutex sync.Mutex
	readyMutex   sync.Mutex
}

func New(logger *slog.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (m *Server) SetHealthz(healthz bool) {
	m.healthzMutex.Lock()
	defer m.healthzMutex.Unlock()
	m.healthz = healthz
}

func (m *Server) SetReady(ready bool) {
	m.readyMutex.Lock()
	defer m.readyMutex.Unlock()
	m.ready = ready
}

func (m *Server) IsHealthy() bool {
	m.healthzMutex.Lock()
	defer m.healthzMutex.Unlock()
	return m.healthz
}

func (m *Server) IsReady() bool {
	m.readyMutex.Lock()
	defer m.readyMutex.Unlock()
	return m.ready
}

func (m *Server) readyHandler(w http.ResponseWriter, _ *http.Request) {
	ready := m.IsReady()
	if ready {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "DOWN")
	}
}

func (m *Server) healthzHandler(w http.ResponseWriter, _ *http.Request) {
	healthz := m.IsHealthy()
	if healthz {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintln(w, "DOWN")
	}
}

func (m *Server) Serve(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", m.healthzHandler)
	mux.HandleFunc("/healthz", m.healthzHandler)
	mux.HandleFunc("/ready", m.readyHandler)

	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		m.logger.Error(
			"failed to start metrics server",
			"err", err,
		)
	}
}
