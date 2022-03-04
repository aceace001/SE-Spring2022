package test

import (
	"testing"
)

//login and register
func TestLogin(t *testing.T) {

	var (
		idIn   string
		nameIn string
	)

	idIn = "inputId"
	nameIn = "inputName"

	ts := &TestStruct{}
	idOut, nameOut := ts.StructFunc(idIn, nameIn)

	if idOut == idIn && nameOut == nameIn {
		t.Log("Test passed！")
	} else {
		t.Error("function execution error")
	}

}

//User online broadcast
//Modify username
//Online user inquiry
//public chat
//private chat
//timeout force offline
