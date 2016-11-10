package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

// NewRESTClient returns an initialized RESTClient with some sane
// defaults for timeouts
func NewRESTClient(apiKey string) *RESTClient {
	shortTimeout := 10 * time.Second
	return &RESTClient{
		APIKey: apiKey,
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				Dial:                (&net.Dialer{Timeout: shortTimeout}).Dial,
				TLSHandshakeTimeout: shortTimeout}}}
}

// RESTClient wraps an http.Client and provides helper methods
// to make RESTful interaction a little easier
type RESTClient struct {
	*http.Client
	APIKey           string
	GetResourceFn    func(string) ([]byte, error)
	PostResourceFn   func(string, interface{}) ([]byte, error)
	DeleteResourceFn func(string) ([]byte, error)
	Trace            bool
}

// GetResource returns raw byte slice from a REST API endpoint
func (r *RESTClient) GetResource(url string) ([]byte, error) {
	if r.GetResourceFn != nil {
		return r.GetResourceFn(url)
	}
	return r.get(url)
}

// PostResource sends a POST request, and returns the raw response
func (r *RESTClient) PostResource(url string, body interface{}) ([]byte, error) {
	if r.DeleteResourceFn != nil {
		return r.PostResourceFn(url, body)
	}
	return r.post(url, body)
}

// DeleteResource sends a DELETE request, and returns the raw response
func (r *RESTClient) DeleteResource(url string) ([]byte, error) {
	if r.DeleteResourceFn != nil {
		return r.DeleteResourceFn(url)
	}
	return r.delete(url)
}

func (r *RESTClient) get(url string) ([]byte, error) {
	return r.send("GET", url, nil)
}

func (r *RESTClient) post(url string, body interface{}) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return r.send("POST", url, bytes.NewBuffer(data))
}

func (r *RESTClient) delete(url string) ([]byte, error) {
	return r.send("DELETE", url, nil)
}

func (r *RESTClient) send(method string, url string, bodySource io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, bodySource)
	if err != nil {
		return nil, err
	}
	if r.APIKey != "" {
		req.Header["Authorization"] = []string{"Bearer " + r.APIKey}
	}
	if bodySource != nil {
		req.Header["Content-Type"] = []string{"application/json;charset=utf8"}
	}
	if r.Trace {
		r.debug("REQUEST")(httputil.DumpRequestOut(req, true))
	}
	res, err := r.Do(req)
	if err != nil {
		return nil, err
	}
	if r.Trace {
		r.debug("RESPONSE")(httputil.DumpResponse(res, true))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return body, nil
	default:
		return body, fmt.Errorf(http.StatusText(res.StatusCode))
	}
}

func (r *RESTClient) debug(label string) func([]byte, error) {
	return func(data []byte, err error) {
		if err == nil {
			log.Printf("%s:\n%s\n", label, string(data))
		} else {
			log.Printf("%s: error %v\n", label, err)
		}
	}
}
