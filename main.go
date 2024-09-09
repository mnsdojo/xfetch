package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/mnsdojo/xfetch/internal/info"
	"github.com/mnsdojo/xfetch/internal/ui"
)

func main() {
	hostname, _ := os.Hostname()
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME") // For Windows
	}

	// Print username@hostname
	color.Cyan.Printf("%s@%s\n\n", username, hostname)

	fmt.Println("┌─────────────────── System Information ───────────────────┐")

	ui.PrintInfo("OS", info.GetOSInfo(), color.FgRed)
	ui.PrintInfo("Machine", info.GetMachineInfo(), color.FgGreen)
	ui.PrintInfo("Kernel", info.GetKernelInfo(), color.FgMagenta)
	ui.PrintInfo("Uptime", info.GetUptime(), color.FgRed)
	ui.PrintInfo("Battery", info.GetBatteryInfo(), color.FgGreen)
	ui.PrintInfo("Shell", info.GetShell(), color.Cyan)
	ui.PrintInfo("CPU", info.GetCPUInfo(), color.Yellow)
	ui.PrintInfo("Memory", info.GetMemoryInfo(), color.FgMagenta)
	ui.PrintInfo("Disk", info.GetDiskInfo(), color.FgGray)

	fmt.Println("└────────────────────────────────────────────────────────────┘")

	ui.PrintColorDots()
}
