package goutils

import (
	"testing"
	"time"
)

func TestID(t *testing.T) {
	idGen := IDGenSingleton()
	id := idGen.NextID()
	id2 := ((idGen.lastTimestamp - idGen.Epoch) << (idGen.MachineBit + idGen.SequenceBit)) + int64(idGen.MachineId<<idGen.SequenceBit) + 0
	if id <= 0 {
		t.Error("generate id error")
	}
	if id != id2 {
		t.Error("generate id error")
	}
}

func TestIDTillNextMillisecond(t *testing.T) {
	idGen := IDGenSingleton()

	// wait 1 millisecond
	ms := time.Now().UnixNano() / 1000000
	nextMS := idGen.tillNextMillisecond(ms)
	if nextMS != ms+1 {
		t.Error("failed to till next millisecond")
	}

	// wait 10 milliseconds
	ms += 10
	nextMS = idGen.tillNextMillisecond(ms)
	if nextMS != ms+1 {
		t.Error("failed to till next millisecond")
	}
}