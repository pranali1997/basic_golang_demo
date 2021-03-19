package router

import (
	"handsongo/pkg/hello"
	"net/http"
	"github.com/gorilla/mux"
)

func Router() http.Handler{


	// http.HandleFunc("/hello", hello.Hello)
	// http.HandleFunc("/headers", hello.Headers)
	r := mux.NewRouter()
    r.HandleFunc("/hello", hello.Hello)
    r.HandleFunc("/headers", hello.Headers)
    // http.Handle("/", r)

	return r;
	// http.ListenAndServe(":8090", nil)
}
