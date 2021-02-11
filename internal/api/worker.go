package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type workerItem struct {
	index   int
	url     string
	timeout time.Duration

	response ResponseItem
}

func worker(ctx context.Context, parallel int, in <-chan workerItem) <-chan workerItem {
	out := make(chan workerItem)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < parallel; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for r := range in {
					resp := doQuery(ctx, r)
					if resp.response.Error != nil {
						resp.response.Error = fmt.Sprintf("%v", resp.response.Error)
					}
					out <- resp
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func doQuery(ctx context.Context, in workerItem) workerItem {
	result := in

	if ctx.Err() != nil {
		result.response.Error = ctx.Err()
		return result
	}

	ctx, cancel := context.WithTimeout(ctx, in.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, in.url, nil)
	if err != nil {
		result.response.Error = err
		return result
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		result.response.Error = err
		return result
	}

	result.response.Code = resp.StatusCode
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		result.response.Error = fmt.Errorf("bad status code %v", resp.StatusCode)
		return result
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.response.Error = err
		return result
	}

	if !json.Valid(body) {
		result.response.Error = errors.New("not a json")
		return result
	}

	result.response.Response = body

	return result
}
