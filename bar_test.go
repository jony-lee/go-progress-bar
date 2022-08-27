package progress

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	bar := New(100, WithFiller("⭐️"))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 10)
		bar.Done(1)
	}
	bar.Finish()
}
