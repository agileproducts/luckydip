package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETLuckyDip(t *testing.T) {
	dataset := []string{"one","two","three","four","five"}
	server := &luckydipInstance{sourceData:&dataset}
	request, _ := http.NewRequest(http.MethodGet, "/luckydip/3", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	// can't quite work out how to test this
	fmt.Println(response.Body.String())

	/*actual := response.Body.String()
	expected := "cat"

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}*/
}