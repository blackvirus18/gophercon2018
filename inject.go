n := negroni.New(negroni.NewRecovery())
n.Use(instrumentation.StatsDMiddlewareLogger())

