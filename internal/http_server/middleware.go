package http_server

import (
    "http-mux/internal/config"
    "net/http"
    "sync/atomic"
)

func throttler(next http.Handler) http.Handler {
    var connections int32
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        atomic.AddInt32(&connections, 1)
        defer atomic.AddInt32(&connections, -1)

        if atomic.LoadInt32(&connections) >= int32(config.MaxInputRequests) {
            w.WriteHeader(http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}
