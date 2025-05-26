package profilers

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProfileScope100ms(t *testing.T) {

	defer func() {
		// fix due to issue with time resolution of 'time' package constants being 1000 larger than that of the
		// profiler's clock.
		const adjusted_time = int64(time.Millisecond * 100 / 1000)
		var total int64
		for _, scope := range Scopes {
			total += scope.end - scope.start
		}
		total /= int64(len(Scopes))
		if total < adjusted_time {
			assert.Failf(t, "average scope time is lower than the actual execution time", "avg time: %d, expected: %d", total, time.Millisecond*100)
		}

		if total >= adjusted_time+adjusted_time/10 {
			assert.Fail(t, "average scope time is 10 percent or more higher than actual execution time")
		}
	}()

	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 100)
	}()

}

func TestProfileScope10ms(t *testing.T) {

	defer func() {
		// fix due to issue with time resolution of 'time' package constants being 1000 larger than that of the
		// profiler's clock.
		const adjusted_time = int64(time.Millisecond * 10 / 1000)
		var total int64
		for _, scope := range Scopes {
			total += scope.end - scope.start
		}
		total /= int64(len(Scopes))
		if total < adjusted_time {
			assert.Failf(t, "average scope time is lower than the actual execution time", "avg time: %d, expected: %d", total, time.Millisecond*10)
		}

		// if total >= adjusted_time+adjusted_time/10 {
		// 	assert.Fail(t, "average scope time is 10 percent or more higher than actual execution time", "avg: %d, expected < %d", total, adjusted_time+adjusted_time/20)
		// }
	}()

	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()
	defer func() {
		defer ProfileScope("test")()
		time.Sleep(time.Millisecond * 10)
	}()

}
