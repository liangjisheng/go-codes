package gopsutil__test

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	//"github.com/shirou/gopsutil/mem"  // to use v2
	"testing"
)

func TestGoPsUtil(t *testing.T) {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	t.Logf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	t.Log(v)

	c, _ := cpu.Info()
	t.Log(c)

	d, _ := disk.Usage("/Users/liangjisheng")
	t.Log(d)
}
