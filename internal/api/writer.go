package api

import (
	"http-mux/internal/config"
)

func writer(urls []string) <-chan workerItem {
	out := make(chan workerItem, 0)

	go func() {
		for i, v := range urls {
			out <- workerItem{
				index:   i,
				url:     v,
				timeout: config.RequestTimeout,
			}
		}
		close(out)
	}()

	return out
}
