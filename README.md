# go-progress-bar
a progress-bar with golang


### install
```shell
go get -u github.com/jony-lee/go-progress-bar
```
### how to use
```go
package main

import (
	"time"

	progress "github.com/jony-lee/go-progress-bar"
)

func main() {
	bar := progress.New(100)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 10)
		bar.Done(1)
	}
	bar.Finish()
}
```