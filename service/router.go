package service

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	log "github.ibm.com/AdvancedAnalyticsCanada/service-common-lib/common/logging"
)

var (
	g_router = mux.NewRouter()
)

func RegisterHandlerFunction(path string, method string, f func(http.ResponseWriter, *http.Request)) {
	g_router.Path(path).Methods(method).HandlerFunc(f)
}
func RegisterHandler(path string, method string, f http.Handler) {
	g_router.Path(path).Methods(method).Handler(f)
}

// Endpoints router function
//@params
//	Port number to bind
//@return
//
func Router(port string) *mux.Router {

	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Info("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Info("Readyz probe is positive.")
	}()

	g_router.HandleFunc("/readyz", readyz(isReady))
	g_router.Path("/healthz").Methods("GET").HandlerFunc(healthz)
	// TODO: implement metrics endpoint
	g_router.Path("/metrics").Methods("GET").HandlerFunc(healthz)

	n := negroni.Classic()
	n.UseHandler(g_router)
	n.Run(fmt.Sprintf(":%s", port))

	return g_router
}
