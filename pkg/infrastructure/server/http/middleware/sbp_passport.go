package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/config"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/constants"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/utilities"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/request"
)

type sessionSupplier struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"active"`
}

type sessionUserSupplier struct {
	Supplier sessionSupplier `json:"supplier"`
}

type sessionUser struct {
	ID            string                `json:"id"`
	Fullname      string                `json:"fullname"`
	UserSuppliers []sessionUserSupplier `json:"userSuppliers"`
}

type responseSessionData struct {
	TokenSBP string      `json:"tokenSBP"`
	User     sessionUser `json:"user"`
}

// ResponseSession ...
type ResponseSession struct {
	Message string              `json:"message"`
	Data    responseSessionData `json:"data"`
}

// SbpPassport used for handling token header
func SbpPassport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := config.NewConfig()

		host := fmt.Sprintf("%s/verification-ssc-token", conf.SessionAPIHost)
		utilities.Log("%s", host)

		client := http.Client{}
		req, _ := http.NewRequest("GET", host, nil)

		token := utilities.ExtractBearerToken(r)

		if token == "" {
			utilities.JSONError(w, r, "UNAUTHORIZED", 400, nil)
			return
		}

		// set the request header bearer
		req.Header.Set("authorization", "Bearer "+token)

		utilities.Log("Requesting to %s", host)
		resp, err := client.Do(req)
		if err != nil || (resp.StatusCode > http.StatusOK && resp.StatusCode < http.StatusBadRequest) {
			utilities.JSONError(w, r, "UNAUTHORIZED", 400, nil)
			return
		}

		rs := ResponseSession{}
		err = json.NewDecoder(resp.Body).Decode(&rs)
		if err != nil {
			utilities.JSONError(w, r, "ERR_PARSE_BODY", 400, nil)
			return
		}

		// extract user id
		userID, err := strconv.Atoi(rs.Data.User.ID)
		if err != nil {
			utilities.JSONError(w, r, "ERR_PARSE_BODY", 400, nil)
			return
		}

		// extract supplier id
		supplierID, err := strconv.Atoi(rs.Data.User.UserSuppliers[0].Supplier.ID)
		if err != nil {
			utilities.JSONError(w, r, "ERR_PARSE_BODY", 400, nil)
			return
		}

		ri := request.Info{
			UserID:     userID,
			UserName:   rs.Data.User.Fullname,
			SupplierID: supplierID,
			TokenSBP:   rs.Data.TokenSBP,
		}

		ctx := context.WithValue(r.Context(), constants.CONTEXT_USER, &ri)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
