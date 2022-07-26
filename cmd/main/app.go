package main

import (
	"awesomeProject/internal/user"
	"awesomeProject/package/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	logger := logging.GetLogger()
	logger.Info("create router") //info -
	router := httprouter.New()

	logger.Info("register user handler")
	//router.GET("/:name", IndexHandler)
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info("server is listening")
	logger.Info(server.Serve(listener))
}
