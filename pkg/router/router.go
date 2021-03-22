package router

import (
	"handsongo/pkg/hello"
	"net/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func NewRouter(lgr *zap.Logger,) http.Handler{


	// http.HandleFunc("/hello", hello.Hello)
	// http.HandleFunc("/headers", hello.Headers)
	router := mux.NewRouter()
    router.HandleFunc("/hello", hello.Hello)
	router.HandleFunc("/hello/error", hello.HelloError)
    router.HandleFunc("/headers", hello.Headers)
    // http.Handle("/", r)

	return router;
	// http.ListenAndServe(":8090", nil)
}
