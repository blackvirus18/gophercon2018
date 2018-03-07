func IncrementInStatsD(key string) bool {
if statsD == nil {
return false
}
statsD.Increment(key)
return true
}
