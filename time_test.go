package goutils

import (
	"testing"
)

func TestUnixTime(t *testing.T) {
	unixTime := UnixTime()
	unixMilliTime := UnixMilliTime()
	unixMicroTime := UnixMicroTime()
	unixNanoTime := UnixNanoTime()

	if unixNanoTime/unixMicroTime != 1000 {
		t.Error("UnixNanoTime or UnixMicroTime error")
	}

	if unixMicroTime/unixMilliTime != 1000 {
		t.Error("UnixMicroTime or UnixMilliTime error")
	}

	if unixMilliTime/unixTime != 1000 {
		t.Error("UnixMilliTime or UnixTime error")
	}

	if unixTime < 0 {
		t.Error("or UnixTime error")
	}
}
