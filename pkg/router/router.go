package router

import (
	"handsongo/pkg/hello"
	"net/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"handsongo/pkg/user"
)

func NewRouter(lgr *zap.Logger,userService user.Service) http.Handler{


	// http.HandleFunc("/hello", hello.Hello)
	// http.HandleFunc("/headers", hello.Headers)
	router := mux.NewRouter()
	uh := user.NewUserHandler(userService)

    router.HandleFunc("/hello", hello.Hello)
	router.HandleFunc("/hello/error", hello.HelloError)
    router.HandleFunc("/headers", hello.Headers)
	router.HandleFunc("/user", uh.CreateUser).Methods(http.MethodPost)

    // http.Handle("/", r)

	return router;
	// http.ListenAndServe(":8090", nil)
}
