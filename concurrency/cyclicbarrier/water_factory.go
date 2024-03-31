package cyclicbarrier

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

//工厂提供多条生产线，每条负责生产氧原子（N 条）或氢原子（2N 条），各由一个 goroutine 负责
//通过一个栅栏，只有一个氧原子和两个氢原子准备好，才能生成出一个水分子，否则所有生产线都处于等待状态
//水分子是逐个按照顺序产生的（原子种类和数量有要求）

//需要引入
//信号量 semaH: 控制氢原子。空槽数资源数设置为 2
//信号量 semaO: 控制氧原子。空槽数资源数设置为 1
//循环栅栏：等待两个氢原子和一个氧原子填补空槽，直到任务完成

type H2O struct {
	semaH *semaphore.Weighted
	semaO *semaphore.Weighted
	b     cyclicbarrier.CyclicBarrier
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2),
		semaO: semaphore.NewWeighted(1),
		b:     cyclicbarrier.New(3),
	}
}

func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)
	releaseHydrogen()                 // 输出 H
	h2o.b.Await(context.Background()) // 等待栅栏放行
	h2o.semaH.Release(1)              // 释放氢原子空槽
}

func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	releaseOxygen()                   // 输出 O
	h2o.b.Await(context.Background()) // 等待栅栏放行
	h2o.semaO.Release(1)              // 释放氢原子空槽
}
