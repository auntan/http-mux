package api

import (
	"context"
	"encoding/json"

	"http-mux/internal/config"
)

type Response struct {
	Responses []ResponseItem `json:",omitempty"`
	Error     interface{}    `json:",omitempty"`
}

type ResponseItem struct {
	Code     int             `json:",omitempty"`
	Response json.RawMessage `json:",omitempty"`
	Error    interface{}     `json:",omitempty"`
}

func QueryUrls(ctx context.Context, urls []string) Response {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	toWorker := writer(urls)
	responses := worker(ctx, config.ParallelRequests, toWorker)

	result := Response{
		Responses: make([]ResponseItem, len(urls)),
	}

	for resp := range responses {
		result.Responses[resp.index] = resp.response
		if resp.response.Error != nil {
			result.Error = "internal error, check details"
			cancel()
		}
	}

	return result
}
