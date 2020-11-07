package snowflake

import (
	"sync"
	"time"
	"errors"
)

//雪花算法
//41bit timestamp | 10 bit machineID : 5bit dataCenterID 5bit workerID ｜ 12 bit sequenceBits
//最多使用69年

const (
	workerIDBits =  uint64(5)  		// 10bit 工作机器ID中的 5bit workerID
	dataCenterIDBits = uint64(5) 	// 10 bit 工作机器ID中的 5bit dataCenterID
	sequenceBits = uint64(12)

	maxWorkerID = int64(-1) ^ (int64(-1) << workerIDBits) //节点ID的最大值 用于防止溢出
	maxDataCenterID = int64(-1) ^ (int64(-1) << dataCenterIDBits)
	maxSequence = int64(-1) ^ (int64(-1) << sequenceBits)

	timeLeft = uint8(22)  // timeLeft = workerIDBits + dataCenterIDBits + sequenceBits // 时间戳向左偏移量
	dataLeft = uint8(17)  // dataLeft = workerIDBits + sequenceBits
	workLeft = uint8(12)  // workLeft = sequenceBits // 节点IDx向左偏移量
	// 2020-11-07 15:00:00 CST
	twepoch = int64(1604732400000) // 常量时间戳(毫秒)
)

// Worker 定义 Worker 工作节点
type Worker struct {
	mu sync.Mutex
	LastStamp int64 // 记录上一次 ID 的时间截
	WorkerID int64 // 该节点的 ID
	DataCenterID int64 // 该节点的数据中心 ID
	Sequence int64 // 当前毫秒已经生成的ID序列号(从0 开始累加) 1毫秒内最多生成4096个ID
}

// NewWorker 分布式情况下,我们应通过外部配置文件或其他方式为每台机器分配独立的id
func NewWorker(workerID,dataCenterID int64) *Worker  {
	return &Worker{
		WorkerID: workerID,
		LastStamp: 0,
		Sequence: 0,
		DataCenterID: dataCenterID,
	}
}

func (w *Worker) getMilliSeconds() int64 {
	return time.Now().UnixNano() / 1e6
}

func (w *Worker)NextID() (uint64,error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.nextID()
}

func (w *Worker)nextID() (uint64,error) {
	timeStamp := w.getMilliSeconds()
	// 确保当前时间截大于上一次生成 ID 的时间截, 否则会出现重复
	if timeStamp < w.LastStamp{
		return 0,errors.New("time is moving backwards,waiting until")
	}

	if w.LastStamp == timeStamp{
		w.Sequence = (w.Sequence + 1) & maxSequence
		// 如果当前毫秒已经生成的id序列号溢出了, 则需要等待下一毫秒, 如果不等待, 就会导致很多重复
		if w.Sequence == 0 {
			for timeStamp <= w.LastStamp {
				timeStamp = w.getMilliSeconds()
			}
		}
	}else {
		//如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.Sequence = 0
	}

	w.LastStamp = timeStamp
	id := ((timeStamp - twepoch) << timeLeft) |
		(w.DataCenterID << dataLeft)  |
		(w.WorkerID << workLeft) |
		w.Sequence

	return uint64(id),nil
}
