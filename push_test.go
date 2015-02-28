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

	message := Message{"hello"}

	status := endPoint.Send(message)
	compare(t, http.StatusOK, status)
}

func TestSendInvalidUserKey(t *testing.T) {
	server := newTestServer(http.StatusUnauthorized)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	message := Message{"hello"}

	status := endPoint.Send(message)
	compare(t, http.StatusUnauthorized, status)
}
