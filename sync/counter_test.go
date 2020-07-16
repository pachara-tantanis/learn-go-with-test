package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(t, counter)
	})
}

func assertCount(t *testing.T, counter Counter) {
	t.Helper()
	if counter.Value() != 3 {
		t.Errorf("got %d, want %d", counter.Value(), 3)
	}
}