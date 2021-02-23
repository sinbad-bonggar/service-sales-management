package utilities

import "testing"

func TestArrayIntToString(t *testing.T) {
	arrNum := []int{1, 2, 3, 4}

	want := "1,2,3,4"

	got := ArrayIntToString(arrNum, ",")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestStringToArrayInt(t *testing.T) {
	s := "1,2,3,4"

	vals, err := StringToArrayInt(s, ",")
	if err != nil {
		t.Errorf("got error %v", err)
	}

	if len(vals) != 4 {
		t.Errorf("got %d want %d", len(vals), 4)
	}

	// with other delim
	s = "1|2|3|4"

	vals, err = StringToArrayInt(s, "|")
	if err != nil {
		t.Errorf("got error %v", err)
	}

	if len(vals) != 4 {
		t.Errorf("got %d want %d", len(vals), 4)
	}

	// with space
	s = "1,2,3, 4"

	vals, err = StringToArrayInt(s, ",")
	if err == nil {
		t.Errorf("should be error (%s)", s)
	}
}
