package main 

import (
	"testing"
)

// test register
func TestRegister(t *testing.T) {

}
// test login
func TestLogin(t *testing.T) {

}

// test logout
func TestDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "4")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteEntry)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
