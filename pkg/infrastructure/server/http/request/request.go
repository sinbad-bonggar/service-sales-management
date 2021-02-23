package request

import (
	"errors"
	"net/http"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/constants"
)

// Info ...
type Info struct {
	UserID     int
	SupplierID int
	TokenSBP   string
	UserName   string
}

// TokenInfoFromContext ...
func TokenInfoFromContext(r *http.Request) (*Info, error) {
	val := r.Context().Value(constants.CONTEXT_USER)

	if val == nil {
		return nil, errors.New("INVALID_CONTEXT")
	}

	return val.(*Info), nil
}

// SearchPagination ...
type SearchPagination struct {
	Page   int
	Length int
	Search string
}

// SearchPaginationFromContext ...
func SearchPaginationFromContext(r *http.Request) *SearchPagination {
	val := r.Context().Value(constants.CONTEXT_SEARCH_PAGINATION)

	if val == nil {
		return nil
	}

	return val.(*SearchPagination)
}
