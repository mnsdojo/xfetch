package info

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func GetOSInfo() string {
	hostInfo, _ := host.Info()
	return fmt.Sprintf("%s %s %s", hostInfo.Platform, hostInfo.PlatformVersion, runtime.GOARCH)
}

func GetMachineInfo() string {
	hostInfo, _ := host.Info()
	return hostInfo.Hostname
}

func GetKernelInfo() string {
	hostInfo, _ := host.Info()
	return hostInfo.KernelVersion
}

func GetUptime() string {
	hostInfo, _ := host.Info()
	uptime := time.Duration(hostInfo.Uptime) * time.Second
	return uptime.Round(time.Minute).String()
}

func GetShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "Unknown"
	}
	return shell
}

func GetCPUInfo() string {
	cpuInfo, _ := cpu.Info()
	if len(cpuInfo) > 0 {
		return fmt.Sprintf("%s (%d) @ %.2f GHz", cpuInfo[0].ModelName, runtime.NumCPU(), cpuInfo[0].Mhz/1000)
	}
	return "Unknown"
}

func GetMemoryInfo() string {
	v, _ := mem.VirtualMemory()
	return fmt.Sprintf("%.2f GiB / %.2f GiB (%.0f%%)", float64(v.Used)/(1024*1024*1024), float64(v.Total)/(1024*1024*1024), v.UsedPercent)
}
