package api

import (
	"fmt"
	"net/http"
)


func ServerUp() {
	http.Handle("/", CustomHandler{"Hello, my lil dear WITH PORT"})

	http.ListenAndServe(":8080", nil)
}

type CustomHandler struct {
	message string
}
func (ch CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, ch.message)
}
