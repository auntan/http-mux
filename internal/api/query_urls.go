package api

import (
	"context"
	"encoding/json"
	"fmt"
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
	if err := validate(urls); err != nil {
		return Response{
			Error: fmt.Sprintf("%v", err),
		}
	}

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

func validate(urls []string) error {
	if len(urls) > config.MaxUrls {
		return fmt.Errorf("too much urls, %v allowed but %v provided", config.MaxUrls, len(urls))
	}
	return nil
}
