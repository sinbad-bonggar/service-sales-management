package utilities

// LimitOffset ...
type LimitOffset struct {
	Limit  int
	Offset int
}

// PaginationToOffsetLimit ...
func PaginationToOffsetLimit(page, length int) LimitOffset {
	if page < 1 {
		page = 1
	}

	return LimitOffset{
		Limit:  length,
		Offset: (page - 1) * length,
	}
}
