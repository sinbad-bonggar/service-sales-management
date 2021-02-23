package utilities

import (
	"encoding/json"
	"net/http"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/constants"
)

const (
	jsonParseFailed = "{\"reason\":\"failed to parse json\"}"
)

// ResponsMeta .
type ResponsMeta struct {
	Total int64 `json:"total"`
	Limit int   `json:"limit"`
	Skip  int   `json:"skip"`
}

// JSONResponse ...
// Status -> 200, 400, 401
// Message -> error message
// Data -> Payload
// Meta -> Pagination etc
type JSONResponse struct {
	Status  int         `json:"status"`
	Type    interface{} `json:"type"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
	Meta    interface{} `json:"meta"`
	Code    interface{} `json:"code"`
}

// JSON ...
func JSON(w http.ResponseWriter, r *http.Request, errorMessage *string, status int, Type, data interface{}, meta interface{}) {
	var payload []byte

	var err error

	resp := JSONResponse{}

	resp = JSONResponse{
		Status:  status,
		Message: errorMessage,
		Type:    Type,
		Code:    nil,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set(constants.HttpHeaderContentType, constants.HttpHeaderApplicationJson)

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if payload, err = json.Marshal(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(jsonParseFailed))
		return
	}

	w.Write(payload)
}

// JSONSuccess ...
func JSONSuccess(w http.ResponseWriter, r *http.Request, data interface{}, meta interface{}) {
	JSON(w, r, nil, http.StatusOK, nil, data, meta)
}

// JSONError ...
func JSONError(w http.ResponseWriter, r *http.Request, message string, status int, Type interface{}) {
	JSON(w, r, &message, status, Type, nil, nil)
}
