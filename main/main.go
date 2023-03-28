package main

import (
	"time"

	"MaricoHan/go-progress-bar"
)

func main() {
	bar := progress.New(100, progress.WithFiller("❤️"), progress.WithInterval(time.Second/5))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()
	bar = progress.New(100, progress.WithFiller("⭐️"), progress.WithInterval(time.Second/5))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()
	bar = progress.New(100, progress.WithFiller("龘"), progress.WithInterval(time.Second/5))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()
	bar = progress.New(100, progress.WithFiller("🍺"), progress.WithInterval(time.Second/5))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

	bar = progress.New(100, progress.WithFiller(">"), progress.WithInterval(time.Second/5))
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

}
