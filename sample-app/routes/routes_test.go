package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_should_return_all_customers_when_there_is_no_error(t *testing.T) {
	r := http.NewServeMux()
	r.HandleFunc("/customers", GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}
