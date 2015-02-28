package gopo

import (
	"encoding/json"
	"net/http"
	"net/url"
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

type Response struct {
	Status  int    `json:"status"`
	Request string `json:"request"`
}

func createResponse(resp *http.Response) *Response {
	gopoResponse := &Response{}
	json.NewDecoder(resp.Body).Decode(gopoResponse)

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
