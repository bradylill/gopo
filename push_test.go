package gopo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func compare(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %+v, got: %+v", expected, actual)
	}
}

func newTestServer(statusCode int, response *Response) *httptest.Server {
	handler := http.HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(statusCode)
			b, _ := json.Marshal(response)
			writer.Write(b)
		})

	return httptest.NewServer(handler)
}

func newTestEndPoint(server *httptest.Server, userKey, apiToken string) *EndPoint {
	return &EndPoint{server.URL, userKey, apiToken}
}

func TestSendSuccess(t *testing.T) {
	testResponse := &Response{1, "reqId123"}
	server := newTestServer(http.StatusOK, testResponse)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Send(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, 1, resp.Status)
	compare(t, "reqId123", resp.Request)
}

func TestSendInvalidSend(t *testing.T) {
	testResponse := &Response{0, "reqId123"}
	server := newTestServer(http.StatusBadRequest, testResponse)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Send(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, 0, resp.Status)
	compare(t, "reqId123", resp.Request)
}
