package api

import (
	"encoding/json"
	"net/http"
)

func (api *API) FetchAllClass(w http.ResponseWriter, r *http.Request) {
	classes, err := api.classService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classes)
}
