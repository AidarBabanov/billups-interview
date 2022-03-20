package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	logs "github.com/sirupsen/logrus"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ContentTypeJson = "application/json"
)

func WriteJson(w http.ResponseWriter, httpCode int, data any) error {
	w.Header().Set(ContentType, ContentTypeJson)
	w.WriteHeader(httpCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(&data)
		if err != nil {
			return fmt.Errorf("can't encode REST response to JSON: %w", err)
		}
	}
	return nil
}

func WriteError(w http.ResponseWriter, err error) {
	var apiError *HTTPError
	if errors.As(err, &apiError) {
	} else {
		apiError = InternalServerErrorf("internal server error")
	}
	if err := WriteJson(w, apiError.Code, apiError); err != nil {
		logs.Error(err)
	}
}

type APIHandler func(w http.ResponseWriter, r *http.Request) error

func (h APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		WriteError(w, err)
	}
}

// APIHandlerFunc gets APIHandler and returns standard http.HandlerFunc
func APIHandlerFunc(fn APIHandler) http.HandlerFunc {
	return fn.ServeHTTP
}
