package gopo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func compare(t *testing.T, message string, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s - Expected: %+v, got: %+v", message, expected, actual)
	}
}

func newTestServer(statusCode int, response *Response) *httptest.Server {
	handler := http.HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Content-Type", "application/json")
			writer.Header().Set("X-Limit-App-Limit", strconv.Itoa(response.Usage.Limit))
			writer.Header().Set("X-Limit-App-Remaining", strconv.Itoa(response.Usage.Remaining))
			writer.Header().Set("X-Limit-App-Reset", response.Usage.NextReset)
			writer.WriteHeader(statusCode)

			b, _ := json.Marshal(response)
			writer.Write(b)
		})

	return httptest.NewServer(handler)
}

func newTestEndPoint(server *httptest.Server, userKey, apiToken string) *EndPoint {
	return &EndPoint{server.URL, userKey, apiToken}
}

func TestPushSuccess(t *testing.T) {
	testResponse := &Response{1, "reqId123", []string{}, Usage{4, 2, "123"}}
	server := newTestServer(http.StatusOK, testResponse)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Push(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, "Response", testResponse, resp)
}

func TestPushFail(t *testing.T) {
	testResponse := &Response{0, "reqId123", []string{"error message"}, Usage{2, 1, "1234"}}
	server := newTestServer(http.StatusBadRequest, testResponse)
	defer server.Close()
	endPoint := newTestEndPoint(server, "userKey", "apiToken")

	resp, err := endPoint.Push(Message{})
	if err != nil {
		t.Fail()
	}

	compare(t, "Response", testResponse, resp)
}
