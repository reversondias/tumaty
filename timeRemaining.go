package main

import "time"

func getTimeRemaining(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
