package gopo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const defaultEndPoint = "https://api.pushover.net/1/messages.json"

type EndPoint struct {
	URL      string
	UserKey  string
	APIToken string
}

type Message struct {
	Message string
}

type Usage struct {
	Limit     int
	Remaining int
	NextReset string
}

type Response struct {
	Status  int      `json:"status"`
	Request string   `json:"request"`
	Errors  []string `json:"errors"`
	Usage   Usage
}

func headerValueAsInt(headers http.Header, headerKey string) int {
	v, err := strconv.Atoi(headers.Get(headerKey))
	if err != nil {
		v = -1
	}

	return v
}

func createResponse(resp *http.Response) *Response {
	gopoResponse := &Response{}
	json.NewDecoder(resp.Body).Decode(gopoResponse)

	gopoResponse.Usage.Limit = headerValueAsInt(resp.Header, "X-Limit-App-Limit")
	gopoResponse.Usage.Remaining = headerValueAsInt(resp.Header, "X-Limit-App-Remaining")
	gopoResponse.Usage.NextReset = resp.Header.Get("X-Limit-App-Reset")

	return gopoResponse
}

func NewGopo(userKey, apiToken string) *EndPoint {
	return &EndPoint{defaultEndPoint, userKey, apiToken}
}

func (e EndPoint) Send(message Message) (*Response, error) {
	vals := url.Values{
		"message": {message.Message},
		"user":    {e.UserKey},
		"token":   {e.APIToken}}

	resp, err := http.PostForm(e.URL, vals)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return createResponse(resp), nil
}
