package internal

import (
	"time"

	"github.com/gosuri/uiprogress"
)

type Progress struct {
	Bar *uiprogress.Bar
}

func NewProgress(prependText string, max int) *Progress {
	bar := uiprogress.AddBar(max).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return prependText
	})
	return &Progress{Bar: bar}
}

func (p *Progress) Start() {
	uiprogress.Start()
}

func (p *Progress) Increment() {
	for p.Bar.Incr() {
		time.Sleep(time.Millisecond * 20)
	}
}

func (p *Progress) Stop() {
	uiprogress.Stop()
}
