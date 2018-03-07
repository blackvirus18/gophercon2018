import "net/http/pprof"

func Router(dependencies *service.Instances) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", raven.RecoveryHandler(h.PingHandler)).Methods("GET")
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	router.HandleFunc("/debug/pprof/goroutine", pprof.Index)
	return router
}
