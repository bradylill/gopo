package gopo

import (
	"net/http"
	"net/url"
)

const defaultEndPoint = "https://api.pushover.net/1/messages.json"

type EndPoint struct {
	URL string
}

type Message struct {
	Message string
	User    string
	Token   string
}

func NewGopo() *EndPoint {
	return &EndPoint{defaultEndPoint}
}

func (e EndPoint) Send(message Message) int {
	vals := url.Values{
		"message": {message.Message},
		"user":    {message.User},
		"token":   {message.Token}}

	resp, err := http.PostForm(e.URL, vals)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
