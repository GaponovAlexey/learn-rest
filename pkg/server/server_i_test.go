package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiServer(t *testing.T) {
	s := New(NewConfig())                            //config
	rec := httptest.NewRecorder()                    //config rec
	req, _ := http.NewRequest("GET", "http://", nil) //config req
	s.Hallo().ServeHTTP(rec, req)                    //start server func testing
	assert.Equal(t, rec.Body.String(), "hi")         // enter parameter testing
}
