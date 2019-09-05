package service

import (
	"os"
	"syscall"
	"net/http"
	"os/signal"
	"context"

	log "github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/logging"
	"github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/util"
)

// StartWebServer starts a webserver on the specified port.
//@params
//
//@return
//
func StartServer(port string) {

	log.Info("Service RELEASE: ",  util.GetENV("RELEASE"), " and Port:" , port)
	if port == "" { log.Error("Can't start the server, port number unspecified.") ; return }

	r := Router(port)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	// this channel is for graceful shutdown..
	// if we receive an error, we can send it here to notify the server to be stopped
	shutdown := make(chan struct{}, 1)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			shutdown <- struct{}{}
			log.Error("%v", err)
		}
	}()
	log.Debug("The service is ready to listen and serve.")

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			log.Debug("Got SIGINT...")
		case syscall.SIGTERM:
			log.Debug("Got SIGTERM...")
		}
	case <-shutdown:
		log.Error("Got an error - shutdown...")
	}

	log.Info("The service is shutting down...")
	srv.Shutdown(context.Background())

}
