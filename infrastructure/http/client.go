package http

import (
	"log"
	"net/http"
)

type HttpClient struct {
	Req    *http.Request
	Client *http.Client
	URL    string
	Method string
}

func (h *HttpClient) Init() {
	req, err := http.NewRequest(h.Method, h.URL, nil)
	if err != nil {
		log.Printf("Erro ao criar a requisição: %v", err)
	}

	h.Req = req
}

func (h *HttpClient) SetHeader(name, value string) {
	h.Req.Header.Set(name, value)
}

func (h *HttpClient) Execute() (*http.Response, error) {
	resp, err := h.Client.Do(h.Req)
	if err != nil {
		log.Printf("Erro na chamada da API: %v", err)
		return nil, err
	}
	return resp, nil
}

func NewHttpClient(url, method string) *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
		URL:    url,
		Method: method,
	}
}
