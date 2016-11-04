package infrastructure

import (
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
	APIKey        string
	GetResourceFn func(string) ([]byte, error)
}

// GetResource returns raw byte slice from a REST API endpoint
func (r *RESTClient) GetResource(url string) ([]byte, error) {
	if r.GetResourceFn != nil {
		return r.GetResourceFn(url)
	}
	return r.get(url)
}

func (r *RESTClient) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if r.APIKey != "" {
		req.Header["Authorization"] = []string{"Bearer " + r.APIKey}
	}
	r.debug("REQUEST")(httputil.DumpRequestOut(req, true))
	res, err := r.Do(req)
	if err != nil {
		return nil, err
	}
	r.debug("RESPONSE")(httputil.DumpResponse(res, true))
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
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
