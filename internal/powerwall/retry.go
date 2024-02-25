package powerwall

import "time"

func retry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, sleep*2, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}
