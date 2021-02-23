package utilities

import "testing"

func TestPaginationToOffsetLimit(t *testing.T) {
	page := 0
	length := 10

	want := LimitOffset{
		Limit:  10,
		Offset: 0,
	}

	got := PaginationToOffsetLimit(page, length)
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}

	page = 1
	length = 10

	want = LimitOffset{
		Limit:  10,
		Offset: 0,
	}

	got = PaginationToOffsetLimit(page, length)
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}

	page = 2
	length = 5

	want = LimitOffset{
		Limit:  5,
		Offset: 5,
	}

	got = PaginationToOffsetLimit(page, length)
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}
