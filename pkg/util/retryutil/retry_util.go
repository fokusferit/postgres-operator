package retryutil

import (
	"fmt"
	"time"
)

// Retry calls ConditionFunc until it returns boolean true, a timeout expires or an error occurs.
func Retry(interval time.Duration, timeout time.Duration, f func() (bool, error)) error {
	//TODO: make the retry exponential
	if timeout < interval {
		return fmt.Errorf("timout(%s) should be greater than interval(%v)", timeout, interval)
	}
	maxRetries := int(timeout / interval)
	tick := time.NewTicker(interval)
	defer tick.Stop()

	for i := 0; ; i++ {
		ok, err := f()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		if i+1 == maxRetries {
			break
		}
		<-tick.C
	}
	return fmt.Errorf("still failing after %d retries", maxRetries)
}
