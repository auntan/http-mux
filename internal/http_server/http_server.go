package http_server

import (
    "context"
    "http-mux/internal/config"
    "net/http"
    "sync"
    "time"
)

type HTTPServer struct {
    server *http.Server
    stop   sync.WaitGroup
}

func New() *HTTPServer {
    return &HTTPServer{
        server: &http.Server{
            Addr: config.Host,
        },
    }
}

func (s *HTTPServer) Start() error {
    http.Handle("/", throttler(http.HandlerFunc(handler)))
    http.HandleFunc("/test", testHandler)

    s.stop.Add(1)
    if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
        return err
    }

    s.stop.Wait()
    return nil
}

func (s *HTTPServer) Stop() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := s.server.Shutdown(ctx); err != nil {
        // log
    }

    s.stop.Done()
}
