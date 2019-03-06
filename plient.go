package plient

import (
	"net/http"
	"net/url"
)

type Plient struct {
	client  *http.Client
	headers []Header
}

type Header struct {
	key   string
	value string
}

func create(proxy string, headers []Header) *Plient {
	proxyUrl, err := url.Parse(proxy);
	if err != nil {
		panic("Proxy error")
	}

	client := &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}}
	return &Plient{client: client, headers: headers}
}

func (p Plient) Get(_url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return nil, err
	}
	p.prepare(req)
	return p.client.Do(req)
}

func (p Plient) prepare(req *http.Request) {
	for _, header := range p.headers {
		req.Header.Set(header.key, header.value)
	}
}
