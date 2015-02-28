package gopo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func compare(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %+v, got: %+v", expected, actual)
	}
}

func newTestServer(status int) *httptest.Server {
	handler := http.HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(status)
		})

	return httptest.NewServer(handler)
}

func newTestEndPoint(server *httptest.Server, userKey, apiToken string) *EndPoint {
	return &EndPoint{server.URL, userKey, apiToken}
}

func TestSendSuccess(t *testing.T) {
	server := newTestServer(http.StatusOK)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Send(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, http.StatusOK, resp.StatusCode)
}

func TestSendInvalidSend(t *testing.T) {
	server := newTestServer(http.StatusBadRequest)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Send(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, http.StatusBadRequest, resp.StatusCode)
}
