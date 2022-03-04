package test

import (
	"testing"
)

// //login and register
// func TestLogin(t *testing.T) {

// 	var (
// 		idIn   string
// 		nameIn string
// 	)

// 	idIn = "inputId"
// 	nameIn = "inputName"

// 	ts := &TestStruct{}
// 	idOut, nameOut := ts.StructFunc(idIn, nameIn)

// 	if idOut == idIn && nameOut == nameIn {
// 		t.Log("Test passed！")
// 	} else {
// 		t.Error("function execution error")
// 	}

// }

//test register
func TestRegister(t *testing.T) {

	var jsonStr = []byte(`{"id":4,"first_name":"xyz","last_name":"pqr","email":"xyz@pqr.com","password":"1234567890"}`)

	req, err := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonStr)) //router.POST("/api/register", users.Register)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Reister)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//User online broadcast
//Modify username
//Online user inquiry
//public chat
//private chat
//timeout force offline
