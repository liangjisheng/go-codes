package cron

import (
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

var (
	once   sync.Once
	zkCron *ZKCron
)

type ZKCron struct {
	cron *cron.Cron
}

func Instance() *ZKCron {
	once.Do(func() {
		zkCron = &ZKCron{
			cron: cron.New(
				cron.WithSeconds(),
				cron.WithChain(cron.SkipIfStillRunning(&CLog{}), cron.Recover(&CLog{})),
			),
		}
		zkCron.cron.Start()
	})
	return zkCron
}

// EverySecond 每秒执行
func (c *ZKCron) EverySecond(f func()) {
	_, _ = c.cron.AddFunc("@every 1s", f)
}

// EveryMinute 每分钟执行
func (c *ZKCron) EveryMinute(f func()) {
	_, _ = c.cron.AddFunc("@every 1m", func() {
		f()
	})
}

// EveryHour 每小时执行
func (c *ZKCron) EveryHour(f func()) {
	_, _ = c.cron.AddFunc("@hourly", func() {
		f()
	})
}

func (c *ZKCron) AddFunc(spec string, cmd func()) {
	_, _ = c.cron.AddFunc(spec, cmd)
}

type CLog struct{}

func (l *CLog) Info(msg string, keysAndValues ...interface{}) {
	log.Printf("cron: msg: %s date: %v", msg, keysAndValues)
}

func (l *CLog) Error(err error, msg string, keysAndValues ...interface{}) {
	log.Printf("cron error: %s msg: %s data: %v", err.Error(), msg, keysAndValues)
}
