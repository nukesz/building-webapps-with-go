package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GET_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo?name=Norbi", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	helloWorld(res, req)

	exp := "Hello Norbi!"
	actualResponse := res.Body.String()
	if exp != actualResponse {
		t.Fatalf("Expected %s, got %s", exp, actualResponse)
	}
}

func Test_POST_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("POST", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	helloWorld(res, req)

	exp := "Processing POST request"
	actualResponse := res.Body.String()
	if exp != actualResponse {
		t.Fatalf("Expected %s, got %s", exp, actualResponse)
	}
}
