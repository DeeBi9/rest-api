package handlers

import (
	"net/http"
	"strconv"

	"github.com/Deepanshuisjod/rest-api/db"
	"github.com/gorilla/mux"
)

func AssignTask(rw http.ResponseWriter, r *http.Request) {

}

func MiddlewareValidateID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := vars["id"]
		newid, _ := strconv.Atoi(id)

		idexist := db.CheckID(newid)
		if !idexist {
			rw.Write([]byte("ID does not match"))
			return
		}
		next.ServeHTTP(rw, r)
	})
}
