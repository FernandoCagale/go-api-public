package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mainJson = `{"status":"ok"}`
)

func TestMainPublic(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/public", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, mainPublic(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mainJson, rec.Body.String())
	}
}
