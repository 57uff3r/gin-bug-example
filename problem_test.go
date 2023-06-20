package main

import (
	"bytes"
	"gin-bug-example/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDataModelsCreationWithJSONPayloads(t *testing.T) {
	test_gin_app := api.GetApp()

	// Sign up valid user
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/albums", bytes.NewBufferString("{}"))
	test_gin_app.ServeHTTP(w1, req1)
	assert.Equal(t, 422, w1.Code)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/albums2", bytes.NewBufferString("{}"))
	test_gin_app.ServeHTTP(w1, req2)
	assert.Equal(t, 422, w2.Code)
}
