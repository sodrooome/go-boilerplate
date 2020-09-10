package response

import (
	"backend-project/domain"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	if err != nil {
		JsonResponse(w, status, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JsonResponse(w, http.StatusBadRequest, nil)
}

type cacheResponse domain.Cache

// JSON Response for cache
func (w cacheResponse) Write(b []byte) (int, error) {
	status := w.Status()
	if 200 <= status && status <= 299 {
		w.Cache.Set(w.URlPath, b, cache.DefaultExpiration)
	}
	return w.ResponseWriter.Write(b)
}
