func StatsDMiddlewareLogger() negroni.HandlerFunc {
    return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        vars := mux.Vars(r)
        path := r.URL.Path
        for _, v := range vars {
            path = strings.Replace(path, v, "", len(path))
        }
        key := GetKeyStructure(r.URL.Path)
        t := TimingInStatsD()
        noOfGoRoutine := runtime.NumGoroutine()
        next(rw, r)
        SendInStatsD(key+".time", t)
        IncrementInStatsD(key + ".calls")
        GaugeKeyInStatsD(key+".goroutines", noOfGoRoutine)
    })
}
