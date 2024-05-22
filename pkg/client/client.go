/*

depsdev - CLI client for deps.dev API.
Free access to dependencies, licenses, advisories, and other critical health and security signals for open source package versions.


@author: edoardottt, https://www.edoardoottavianelli.it/

@repository: https://github.com/edoardottt/depsdev

@license: https://github.com/edoardottt/depsdev/blob/main/LICENSE

*/

package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	Seconds30             = 30 * time.Second
	MaxIdleConns          = 100
	IdleConnTimeout       = 90 * time.Second
	TLSHandshakeTimeout   = 10 * time.Second
	ExpectContinueTimeout = 1 * time.Second
)

type Client struct {
	baseURL    string
	underlying *http.Client
}

func New(basePath string) *Client {
	c := &Client{
		baseURL: basePath,
		underlying: &http.Client{
			Transport: &transport{
				underlying: http.Transport{
					Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   Seconds30,
						KeepAlive: Seconds30,
					}).DialContext,
					ForceAttemptHTTP2:     true,
					MaxIdleConns:          MaxIdleConns,
					IdleConnTimeout:       IdleConnTimeout,
					TLSHandshakeTimeout:   TLSHandshakeTimeout,
					ExpectContinueTimeout: ExpectContinueTimeout,
				},
			},
		},
	}

	return c
}

func (c *Client) url(path string) string {
	if strings.Contains(path, "://") {
		return path
	}

	return c.baseURL + path
}

func (c *Client) do(method, path string, target interface{}, body io.Reader) error {
	req, err := http.NewRequest(method, c.url(path), body)
	if err != nil {
		return err
	}

	resp, err := c.underlying.Do(req)

	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if target == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) RawGet(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.url(path), nil)
	if err != nil {
		return []byte{}, err
	}

	resp, err := c.underlying.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer func() { _ = resp.Body.Close() }()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}

func (c *Client) Get(path string, target interface{}) error {
	return c.do(http.MethodGet, path, target, nil)
}

func (c *Client) Post(path string, body interface{}, target interface{}) error {
	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}

	return c.do(http.MethodPost, path, target, &buf)
}
