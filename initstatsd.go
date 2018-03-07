InitiateStatsDMetrics(config StatsDConfig)  {
  address := fmt.Sprintf("%s:%s", host, port)
  statsD, err := statsdv2.New(statsdv2.Address(address), statsdv2.Prefix(appname))
}
