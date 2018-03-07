func TimingInStatsD() *statsdv2.Timing {
if statsD == nil {
return nil
}
t := statsD.NewTiming()
return &t
}
