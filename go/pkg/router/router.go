package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter は新しいルーターを作成します
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

// AddRoute はルートを追加します
func AddRoute(r *mux.Router, path string, handler http.HandlerFunc, methods ...string) {
	r.HandleFunc(path, handler).Methods(methods...)
}
