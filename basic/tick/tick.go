package tick

import "time"

func Do(tick <-chan time.Time) {
	<-tick
}
