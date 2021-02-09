package api

import (
    "context"
    "encoding/json"
    "fmt"
    "http-mux/internal/config"
)

type Response struct {
    Responses []ResponseItem
    Error     interface{}
}

type ResponseItem struct {
    Code     int
    Response json.RawMessage
    Error    interface{}
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
