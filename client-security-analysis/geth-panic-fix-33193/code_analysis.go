// Before fix: Panic risk exists

func randomDuration(min, max time.Duration) time.Duration {

if min > max {

panic("min duration must be less than or equal to max duration")

}

return time.Duration(mrand.Int63n(int64(max-min)) + int64(min))

}

// After fix: Handling boundary conditions

func randomDuration(min, max time.Duration) time.Duration {

if min > max {

panic("min duration must be less than or equal to max duration")

}

if min == max {

return min

}

return time.Duration(mrand.Int63n(int64(max-min)) + int64(min))

}