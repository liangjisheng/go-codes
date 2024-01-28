package cron

import (
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	once   sync.Once
	myCron *MyCron
)

type MyCron struct {
	cron *cron.Cron
}

func Instance() *MyCron {
	once.Do(func() {
		myCron = &MyCron{
			cron: cron.New(
				cron.WithSeconds(),
				cron.WithChain(cron.SkipIfStillRunning(&CLog{}), cron.Recover(&CLog{})),
			),
		}
		myCron.cron.Start()
	})
	return myCron
}

// EverySecond 每秒执行
func (c *MyCron) EverySecond(f func()) {
	_, _ = c.cron.AddFunc("@every 1s", f)
}

// EveryMinute 每分钟执行
func (c *MyCron) EveryMinute(f func()) {
	_, _ = c.cron.AddFunc("@every 1m", func() {
		f()
	})
}

// EveryHour 每小时执行
func (c *MyCron) EveryHour(f func()) {
	_, _ = c.cron.AddFunc("@hourly", func() {
		f()
	})
}

func (c *MyCron) AddFunc(spec string, cmd func()) {
	_, _ = c.cron.AddFunc(spec, cmd)
}

type CLog struct{}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	log.Printf("cron: msg: %s date: %v", msg, keysAndValues)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	log.Printf("cron error: %s msg: %s data: %v", err.Error(), msg, keysAndValues)
}
