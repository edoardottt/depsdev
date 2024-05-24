package depsdev

import (
	"context"

	"github.com/edoardottt/depsdev/pkg/client"
)

type batchBody interface {
	SetNextPageToken(string)
}

type batchResponse[T any] interface {
	GetNextPageToken() string
	Items() []T
}

type batchJob[T any] struct {
	Item T
	Err  error
}

type Iterator[T any] struct {
	cIn     <-chan batchJob[T]
	item    T
	err     error
	hasNext bool

	cancel func()
}

func (v *Iterator[T]) Next() bool {
	bj, ok := <-v.cIn
	if !ok {
		return false
	}

	v.err = bj.Err
	v.item = any(bj.Item).(T)

	return true
}

func (v *Iterator[T]) Item() (T, error) {
	return v.item, v.err
}

func (v *Iterator[T]) Close() {
	if v.cancel != nil {
		v.cancel()
	}
}

func getBatch[T any](ctx context.Context, c *client.Client, endpointPath string, body batchBody, response batchResponse[T]) <-chan batchJob[T] {
	cJob := make(chan batchJob[T])

	go func() {
		defer close(cJob)

		for response.GetNextPageToken() != "" {
			job := batchJob[T]{}

			if err := c.Post(endpointPath, body, &response); err != nil {
				job.Err = err
				cJob <- job

				return
			}

			for _, r := range response.Items() {
				select {
				case <-ctx.Done():
					return
				default:
					job.Item = r
					cJob <- job
				}
			}

			body.SetNextPageToken(response.GetNextPageToken())
		}
	}()

	return cJob
}
