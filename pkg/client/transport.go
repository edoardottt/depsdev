/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://edoardottt.com/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package client

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	retry "github.com/avast/retry-go"
)

const (
	Attempts = 3
	Delay    = time.Millisecond * 100
	Error400 = 400
	Error500 = 500
)

type transport struct {
	underlying http.Transport
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var response *http.Response

	options := []retry.Option{
		retry.Delay(Delay),
		retry.DelayType(retry.BackOffDelay),
		retry.Attempts(Attempts),
		retry.LastErrorOnly(true),
	}

	attempt := 0

	if err := retry.Do(func() error {
		r, err := requestForAttempt(req, attempt)
		if err != nil {
			return retry.Unrecoverable(err)
		}

		attempt++

		resp, err := t.underlying.RoundTrip(r)
		if err != nil {
			return err
		}

		if resp.StatusCode >= Error400 {
			defer func() { _ = resp.Body.Close() }()

			apiError := newErrorResponse(resp)

			if resp.StatusCode < Error500 {
				return retry.Unrecoverable(apiError)
			}

			return apiError
		}

		response = resp

		return nil
	}, options...); err != nil {
		return nil, err
	}

	return response, nil
}

// requestForAttempt returns the request to send at the given attempt.
// The body of a request is consumed when it is sent, so retries get a fresh copy of it.
func requestForAttempt(req *http.Request, attempt int) (*http.Request, error) {
	if attempt == 0 || req.GetBody == nil {
		return req, nil
	}

	body, err := req.GetBody()
	if err != nil {
		return nil, err
	}

	r := req.Clone(req.Context())
	r.Body = body

	return r, nil
}

// newErrorResponse builds an ErrorResponse from the body of an error response.
// The API replies with a plain text message, a JSON {"error": "..."} body is supported as well.
func newErrorResponse(resp *http.Response) *ErrorResponse {
	apiError := &ErrorResponse{StatusCode: resp.StatusCode}

	body, err := io.ReadAll(resp.Body)
	if err == nil && json.Unmarshal(body, apiError) == nil && apiError.Message != "" {
		return apiError
	}

	apiError.Message = strings.TrimSpace(string(body))
	if apiError.Message == "" {
		apiError.Message = http.StatusText(resp.StatusCode)
	}

	return apiError
}
