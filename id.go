package goutils

import (
	"sync"
	"time"
)

// ID 返回一个类snowflake算法生成的ID
func ID() int64 {
	return GenerateID()
}

// GenerateID 返回一个类snowflake算法生成的ID
func GenerateID() int64 {
	return IDGenSingleton().NextID()
}

// 定义 IDGenerator 实例变量
var id *IDGenerator
var once sync.Once

// 获取 IDGenerator 单例
func IDGenSingleton() *IDGenerator {
	once.Do(func() {
		id = &IDGenerator{}
		id.init()
	})
	return id
}

// IDGenerator 结构体
type IDGenerator struct {
	mutex         sync.Mutex // 互斥锁
	Epoch         int64      // 时间戳起点
	MachineId     int        // 机器号
	MachineBit    int        // 机器号最大二进制位数
	Sequence      int        // 序列号
	SequenceBit   int        // 序列号最大二进制位数
	lastTimestamp int64      // 上次ID生成的毫秒级时间戳
}

// 初始化 IDGenerator 结构体变量值，创建 IDGenerator 实例时调用的方法，相当于构造函数
func (ID *IDGenerator) init() {
	ID.Epoch = 1500000000000
	ID.MachineId = 0
	ID.MachineBit = 4
	ID.Sequence = 0
	ID.SequenceBit = 10
	ID.lastTimestamp = 0
}

// 生成一个长整型 ID
func (ID *IDGenerator) NextID() int64 {
	// 互斥锁，确保同一时间只能有一个线程进入
	ID.mutex.Lock()

	// 生成毫秒时间戳 & 序列号
	timestamp := ID.timeGen()
	if ID.lastTimestamp > timestamp { // 判断上次时间戳大于当前时间戳，防止时钟回拨
		timestamp = ID.tilNextMillis(ID.lastTimestamp)
	} else if ID.lastTimestamp == timestamp { // 同一时间戳内序列号自增
		ID.Sequence = (ID.Sequence + 1) & ((1 << ID.SequenceBit) - 1)
		if ID.Sequence == 0 { // 自增序列号超过最大值时，等待到下一毫秒
			timestamp = ID.tilNextMillis(ID.lastTimestamp)
		}
	} else { // 上次时间戳小于当前时间戳则序列号回归为0
		ID.Sequence = 0
	}

	sequence := ID.Sequence
	if timestamp > 0 {
		ID.lastTimestamp = timestamp
	}

	// 解除互斥锁
	ID.mutex.Unlock()

	// 返回生成的ID
	if timestamp > 0 {
		return ((timestamp - ID.Epoch) << (ID.MachineBit + ID.SequenceBit)) + int64(ID.MachineId<<ID.SequenceBit) + int64(sequence)
	}
	return 0 // 等待下一毫秒执行次数过多则返回 ID = 0
}

// 获取当前毫秒级时间戳
func (ID *IDGenerator) timeGen() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取下一毫秒的毫秒级时间戳（相对于lastTimestamp）
func (ID *IDGenerator) tilNextMillis(lastTimestamp int64) int64 {
	timestamp := ID.timeGen()
	count := 0
	for lastTimestamp > timestamp {
		count++
		// 只执行100次
		if count > 100 {
			return 0
		}
		time.Sleep(time.Duration(1) * time.Millisecond)
		timestamp = ID.timeGen()
	}
	return timestamp
}
