package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOk", nil)
	if err != nil {
		fmt.Println("Error at CheckStatusOK:", err)
		return
	}

	resReco := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(CheckStatusOk)
	handlerFunc.ServeHTTP(resReco, req)

	statusCode := resReco.Code
	if statusCode != http.StatusOK {
		t.Errorf("Handler returned %v\n", statusCode)
	}

	want := "Fine!"
	if resReco.Body.String() != want {
		t.Errorf("Handler returned %v\n", resReco.Body.String())
	}

}

func TestCheckNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckNotFound", nil)
	if err != nil {
		fmt.Println("Error at CheckNotFound:", err)
		return
	}

	resReco := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(CheckNotFound)
	handlerFunc.ServeHTTP(resReco, req)

	statusCode := resReco.Code
	if statusCode != http.StatusNotFound {
		t.Errorf("Handler returned %v\n", statusCode)
	}

}
