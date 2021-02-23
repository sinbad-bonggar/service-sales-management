package utilities

import "testing"

type exampleStruct struct {
	Name  string ``
	Email string `valid:"email"`
}

func TestValidator(t *testing.T) {
	testData := exampleStruct{
		Name:  "Sinbad",
		Email: "test@sinbad.co.id",
	}

	got := Validator(testData)
	if got != nil {
		t.Errorf("got %v want nil", got)
	}
}
