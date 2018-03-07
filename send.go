func SendInStatsD(key string, t *statsdv2.Timing) bool {
if statsD == nil {
return false
}
t.Send(key)
return true
}
