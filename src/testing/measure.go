package testing

import (
	"time"
)

const numIterations = 2000

func Measure(f func() error) (time.Duration, error) {
	var durations []time.Duration
	for i := 0; i < numIterations; i++ {
		start := time.Now()
		err := f()
		elapsed := time.Since(start)

		if err != nil {
			return 0, err
		}

		durations = append(durations, elapsed)
	}
	return calculateAverageDuration(durations), nil
}

func calculateAverageDuration(durations []time.Duration) time.Duration {
	total := time.Duration(0)
	for _, duration := range durations {
		total += duration
	}
	count := int64(len(durations))
	return time.Duration(int64(total) / count)
}
