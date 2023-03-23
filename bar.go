package progress

import (
	"fmt"
	"strings"
	"time"

	"github.com/jony-lee/go-progress-bar/unicode"
)

type Bar struct {
	total         int64         // total of task
	current       int64         // current status of task
	filler        string        // filler to progress bar
	filler_size   int           // filler size to progress bar
	filler_length int64         // filler
	time_format   string        // time format
	interval      time.Duration // interval to print progress
	begin         time.Time     // start of task
}

// New 新建进度条实例
func New(total int64, opts ...Option) *Bar {
	bar := &Bar{
		total:         total,
		filler:        "█",
		filler_length: 26,
		time_format:   "15:04:05", // 2006-01-02T15:04:05
		interval:      time.Second,
		begin:         time.Now(),
	}
	for _, opt := range opts {
		opt(bar)
	}
	// 处理宽字符
	bar.filler_size = unicode.GetEastAsianWidth([]rune(bar.filler)[0])
	bar.filler_length = bar.filler_length / int64(bar.filler_size)

	// 定时打印
	ticker := time.NewTicker(bar.interval)
	go func() {
		for bar.current < bar.total {
			fmt.Print(bar.get_progress_string())
			<-ticker.C
		}
	}()
	return bar
}

// Done 更新完成进度
func (bar *Bar) Done(i int64) {
	bar.current += i
}

// Finish 完成最后进度条
func (bar *Bar) Finish() {
	bar.current = bar.total
	fmt.Println(bar.get_progress_string())
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

// WithTimeFormat 设置进度条ETA时间格式
func WithTimeFormat(format string) Option {
	return func(bar *Bar) {
		if len(format) != 0 {
			bar.time_format = format
		}
	}
}

// WithFillerLength 设置进度条显示长度
func WithFillerLength(l int64) Option {
	return func(bar *Bar) {
		if l > 0 {
			bar.filler_length = l
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

// get_percent 获取进度百分比,区间0-100
func (bar *Bar) get_percent() int64 {
	return bar.current * 100 / bar.total
}

// get_eta 获取eta时间
func (bar *Bar) get_eta(now time.Time) string {
	eta := (now.Unix() - bar.begin.Unix()) * 100 / (bar.get_percent() + 1)
	return bar.begin.Add(time.Second * time.Duration(eta)).Format(bar.time_format)
}

// get_progress_string 获取打印控制台字符串
func (bar *Bar) get_progress_string() string {
	fills := bar.get_percent() * bar.filler_length / 100
	chunks := make([]string, bar.filler_length, bar.filler_length)
	blank := make([]byte, bar.filler_size, bar.filler_size)
	for i := 0; i < bar.filler_size; i++ {
		blank[i] = ' '
	}
	for i := int64(0); i < bar.filler_length; i++ {
		switch {
		case i < fills:
			chunks[i] = bar.filler
		default:
			chunks[i] = string(blank)
		}
	}
	now := time.Now()
	eta := bar.get_eta(now)
	qps := bar.current / (now.Unix() - bar.begin.Unix() + 1)
	return fmt.Sprintf("\r[%s]%d/%d [eta]%s [qps]%d ", strings.Join(chunks, ""), bar.current, bar.total, eta, qps)
}
