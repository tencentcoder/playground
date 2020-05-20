package tick

import (
	"testing"
	"time"
)

func BenchmarkDo(b *testing.B) {
	tick := time.Tick(time.Second / 1000)
	b.Run("500", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Do(tick)
		}
	})
}
