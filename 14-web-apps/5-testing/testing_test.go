package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	r = mux.NewRouter()
	endpoints(r)
	os.Exit(m.Run())
}

func Test_mainHandler(t *testing.T) {
	data := []int{}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/Name", bytes.NewBuffer(payload))
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)

	req = httptest.NewRequest(http.MethodGet, "/Name", nil)
	req.Header.Add("content-type", "plain/text")

	rr = httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)
}
