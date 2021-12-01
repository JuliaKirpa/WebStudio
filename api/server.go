package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ServerUp(){
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		_, _ = fmt.Fprintln(w, "Hello, Gipernova!")
	})
	_ = http.ListenAndServe(":8080", r)
}
