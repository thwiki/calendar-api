package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aerogo/aero"
	"github.com/stretchr/testify/assert"
)

func TestStaticRoutes(t *testing.T) {
	app := configure(aero.New())
	request := httptest.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)

	assert.Equal(t, response.Code, http.StatusOK)
}
