package progress

import (
	"fmt"
	"strings"
	"time"
)

type Bar struct {
	total        int64         // total of task
	current      int64         // current status of task
	filler       string        // filler to progress bar
	fillerSize   int           // length displayed
	fillerLength int64         // num of filler repetitions
	timeFormat   string        // time format
	interval     time.Duration // interval to print progress
	begin        time.Time     // start of task
}

// New 新建进度条实例
func New(total int64, opts ...Option) *Bar {
	bar := &Bar{
		total:        total,
		filler:       " >",
		fillerLength: 25,
		timeFormat:   "15:04:05", // 2006-01-02T15:04:05
		interval:     time.Second,
		begin:        time.Now(),
	}
	for _, opt := range opts {
		opt(bar)
	}

	// 适配宽字符
	if len(bar.filler) == len([]rune(bar.filler)) {
		bar.fillerSize = len(bar.filler)
	} else {
		bar.fillerSize = 2
	}

	go func() {
		// 定时打印
		ticker := time.NewTicker(bar.interval)
		defer ticker.Stop()
		for bar.current < bar.total {
			fmt.Print(bar.getProgressString())
			<-ticker.C
		}
	}()
	return bar
}

// Done 更新完成进度
func (b *Bar) Done(i int64) {
	b.current += i
}

// Finish 完成最后进度条
func (b *Bar) Finish() {
	b.current = b.total
	fmt.Println(b.getProgressString())
}

type Option func(*Bar)

// WithFiller 设置进度条填充字符
func WithFiller(filler string) Option {
	return func(bar *Bar) {
		if len(bar.filler) != 0 {
			bar.filler = filler
		}
	}
}

// WithETAFormat 设置进度条 ETA 时间格式
func WithETAFormat(format string) Option {
	return func(bar *Bar) {
		if len(format) != 0 {
			bar.timeFormat = format
		}
	}
}

// WithFillerLength 设置进度条显示长度
func WithFillerLength(l int64) Option {
	return func(bar *Bar) {
		if l > 0 {
			bar.fillerLength = l
		}
	}
}

// WithInterval 设置刷新进度条的时间间隔
func WithInterval(t time.Duration) Option {
	return func(bar *Bar) {
		bar.interval = time.Millisecond * 8
		if t > time.Millisecond*8 { // lower than 125HZ
			bar.interval = t
		}
	}
}

// percent 获取进度百分比,区间0-100
func (b *Bar) percent() int64 {
	return b.current * 100 / b.total
}

// eta 获取eta时间
func (b *Bar) eta(now time.Time) string {
	eta := (now.Unix() - b.begin.Unix()) * 100 / (b.percent() + 1)
	return b.begin.Add(time.Second * time.Duration(eta)).Format(b.timeFormat)
}

// getProgressString 获取打印控制台字符串
func (b *Bar) getProgressString() string {
	fills := int(b.percent() * b.fillerLength / 100)
	now := time.Now()
	qps := b.current / (now.Unix() - b.begin.Unix() + 1)
	return fmt.Sprintf("\r[%s]%d/%d [eta]%s [qps]%d ", strings.Repeat(b.filler, fills)+strings.Repeat(" ", (int(b.fillerLength)-fills)*b.fillerSize), b.current, b.total, b.eta(now), qps)
}
