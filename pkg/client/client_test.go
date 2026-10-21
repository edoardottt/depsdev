/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package client_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/edoardottt/depsdev/pkg/client"
	"github.com/stretchr/testify/require"
)

func TestErrorResponse(t *testing.T) {
	tests := []struct {
		name        string
		status      int
		contentType string
		body        string
		message     string
		requests    int32
	}{
		{
			name:        "plain text client error is not retried",
			status:      http.StatusNotFound,
			contentType: "text/plain; charset=utf-8",
			body:        "package not found\n",
			message:     "package not found",
			requests:    1,
		},
		{
			name:        "JSON client error",
			status:      http.StatusBadRequest,
			contentType: "application/json",
			body:        `{"error":"invalid request"}`,
			message:     "invalid request",
			requests:    1,
		},
		{
			name:        "empty server error is retried",
			status:      http.StatusServiceUnavailable,
			contentType: "text/plain; charset=utf-8",
			body:        "",
			message:     http.StatusText(http.StatusServiceUnavailable),
			requests:    client.Attempts,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requests atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				requests.Add(1)
				w.Header().Set("Content-Type", tt.contentType)
				w.WriteHeader(tt.status)
				_, _ = io.WriteString(w, tt.body)
			}))
			defer server.Close()

			var target map[string]any

			err := client.New(server.URL).Get("/path", &target)
			require.Error(t, err)

			var apiError *client.ErrorResponse

			require.ErrorAs(t, err, &apiError)
			require.Equal(t, tt.status, apiError.StatusCode)
			require.Equal(t, tt.message, apiError.Message)
			require.Equal(t, tt.requests, requests.Load())
		})
	}
}

func TestPostRetrySendsBodyAgain(t *testing.T) {
	var (
		requests atomic.Int32
		bodies   = make(chan string, client.Attempts)
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		bodies <- string(body)

		if requests.Add(1) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)

			return
		}

		_, _ = io.WriteString(w, `{"ok":true}`)
	}))
	defer server.Close()

	var target struct {
		OK bool `json:"ok"`
	}

	err := client.New(server.URL).Post("/batch", map[string]string{"pageToken": "abc"}, &target)
	require.NoError(t, err)
	require.True(t, target.OK)
	require.Equal(t, int32(2), requests.Load())

	close(bodies)

	for body := range bodies {
		require.JSONEq(t, `{"pageToken":"abc"}`, body)
	}
}
