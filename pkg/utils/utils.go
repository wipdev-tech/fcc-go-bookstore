package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wipdev-tech/fcc-go-bookstore/pkg/models"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetBookIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
    id, err := strconv.ParseInt(bookId, 0, 0)
    return id, err
}

func MakeResponse(w http.ResponseWriter, b *models.Book) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}
