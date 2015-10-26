package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	APIVersion = 1.20
)

type Client struct {
	endpoint   *url.URL
	HTTPClient *http.Client
	TLSConfig  *tls.Config
	APIVersion float64
}

func NewClient(e string, tlsConfig *tls.Config, t time.Duration) (*Client, error) {
	u, err := url.Parse(e)
	if err != nil {
		return nil, err
	}

	if u.Scheme == "" || u.Scheme == "tcp" {
		if tlsConfig == nil {
			u.Scheme = "http"
		} else {
			u.Scheme = "https"
		}
	}
	timeout := t * time.Second

	httpClient := newHTTPClient(u, tlsConfig, timeout)
	return &Client{u, httpClient, tlsConfig, APIVersion}, nil
}

func (c *Client) doRequest(method string, path string, body []byte, headers map[string]string) ([]byte, error) {
	b := bytes.NewBuffer(body)
	reader, err := c.doStreamRequest(method, path, b, headers)
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) doStreamRequest(method string, path string, in io.Reader, headers map[string]string) (io.ReadCloser, error) {
	if (method == "POST" || method == "PUT") && in == nil {
		in = bytes.NewReader(nil)
	}
	req, err := http.NewRequest(method, c.endpoint.String()+path, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if headers != nil {
		for header, value := range headers {
			req.Header.Add(header, value)
		}
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, errors.New(string(resp.StatusCode))
	}

	return resp.Body, nil
}

func (c *Client) getAPIVersionString() string {
	return fmt.Sprintf("v%.2f", c.APIVersion)
}

func (c *Client) Ping() ([]byte, error) {
	path := fmt.Sprintf("/%s/_ping", c.getAPIVersionString())
	resp, err := c.doRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) Info() ([]byte, error) {
	path := fmt.Sprintf("/%s/info", c.getAPIVersionString())
	resp, err := c.doRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func newHTTPClient(u *url.URL, tlsConfig *tls.Config, timeout time.Duration) *http.Client {
	httpTransport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	switch u.Scheme {
	default:
		httpTransport.Dial = func(proto, addr string) (net.Conn, error) {
			return net.DialTimeout(proto, addr, timeout)
		}
	case "unix":
		socketPath := u.Path
		unixDial := func(proto, addr string) (net.Conn, error) {
			return net.DialTimeout("unix", socketPath, timeout)
		}
		httpTransport.Dial = unixDial
		// Override the main URL object so the HTTP lib won't complain
		u.Scheme = "http"
		u.Host = "unix.sock"
		u.Path = ""
	}
	return &http.Client{Transport: httpTransport}
}
