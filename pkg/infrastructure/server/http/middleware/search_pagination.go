package middleware

import (
	"context"
	"strconv"

	"net/http"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/constants"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/request"
)

// SearchPagination ...
func SearchPagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}

		rsp := request.SearchPagination{
			Page:   1,
			Length: 10,
			Search: "",
		}

		// get query length
		length := r.URL.Query().Get(constants.QUERY_PARAM_LENGTH)

		if length != "" {
			n, err := strconv.Atoi(length)
			if err == nil {
				rsp.Length = n
			}
		}

		// get query page
		page := r.URL.Query().Get(constants.QUERY_PARAM_PAGE)
		if page != "" {
			n, err := strconv.Atoi(page)
			if err == nil {
				rsp.Page = n
			}
		}

		rsp.Search = r.URL.Query().Get(constants.QUERY_PARAM_SEARCH)

		ctx := context.WithValue(r.Context(), constants.CONTEXT_SEARCH_PAGINATION, &rsp)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
