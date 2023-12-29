package timer

import "time"

func Now() time.Time {
	return time.Now().UTC()
}

func NoSeconds() time.Time {
	now := Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		0,
		0,
		now.Location(),
	)
}
func NowNoMinuteAndSeconds() time.Time {
	now := Now()
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		0,
		0,
		0,
		now.Location(),
	)
}
