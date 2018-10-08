package model

import "testing"

func Test_NewUser_Success(t *testing.T) {
	expect := User {
		Id: "",
		Name: "akase",
		Pass: "akase",
	}

	actual := NewUser("akase", "akase")
	if actual != expect {
		t.Fatalf("failed test acutal_name=%#v, actual_pass=%#v", actual.Name, actual.Pass)
	}
}