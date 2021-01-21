package cancelrequest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	assert.Equal(t, data, response.Body.String())
}

func TestCancelServer(t *testing.T) {
	data := "hello, universe"
	store := &SpyStore{response: data}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Millisecond, cancel)
	request = request.WithContext(cancellingCtx)

	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	assert.Equal(t, true, store.cancelled)
}
