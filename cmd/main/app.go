package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/user"
	"awesomeProject/package/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
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

	cfg := config.GetConfig()

	logger.Info("register user handler")
	//router.GET("/:name", IndexHandler)
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
	//запустить на сокете лучше чем на порте, но из докера надо будет выбрасывать сокет
	//инженекс стоит на хостовой или в контейнере
	//ну это если не хочешь занимать
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDr, err := filepath.Abs(filepath.Dir(os.Args[0])) //где мы находимся
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDr, "app.sock")

		logger.Info("listen socket")
		//logger.Debugf("socket path %s", socketPath)
		listener, listenErr = net.Listen("listen unix socket", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}
	if listenErr != nil {
		panic(listenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//logger.Info("server is listening")
	logger.Info(server.Serve(listener))
}
