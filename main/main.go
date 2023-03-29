package main

import (
	"time"

	"MaricoHan/go-progress-bar"
)

func main() {
	opts := []progress.Option{
		progress.WithInterval(time.Second / 5),
		progress.WithETAFormat("2006-01-02 15:04:05"),
		progress.WithFillerLength(25),
	}

	bar := progress.New(100, append(opts, progress.WithFiller("❤️"))...)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

	bar = progress.New(100, append(opts, progress.WithFiller("⭐️"))...)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

	bar = progress.New(100, append(opts, progress.WithFiller("龘"))...)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

	bar = progress.New(100, append(opts, progress.WithFiller("🍺"))...)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

	bar = progress.New(100, append(opts, progress.WithFiller("-->"))...)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 50)
		bar.Done(1)
	}
	bar.Finish()

}
