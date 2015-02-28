package gopo

import (
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

func NewGopo(userKey, apiToken string) *EndPoint {
	return &EndPoint{defaultEndPoint, userKey, apiToken}
}

func (e EndPoint) Send(message Message) int {
	vals := url.Values{
		"message": {message.Message},
		"user":    {e.UserKey},
		"token":   {e.APIToken}}

	resp, err := http.PostForm(e.URL, vals)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
