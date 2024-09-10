package ui

import (
	"fmt"

	"github.com/gookit/color"
)

var icons = map[string]string{
	"OS":      "🖥️",
	"Machine": "🏠",
	"Kernel":  "🐧",
	"Uptime":  "⏳",
	"Shell":   "🐚",
	"Battery": "🔋",
	"CPU":     "🧠",
	"Memory":  "🧩",
	"Disk":    "💾",
	"CpuTemp": "🌡️",
}

func PrintInfo(label, value string, labelColor color.Color) {
	icon := icons[label]
	labelColor.Printf("%-2s %-10s", icon, label)
	fmt.Printf(" →  %s\n", value)
}

func PrintColorDots() {
	colors := []color.Color{color.FgWhite, color.FgWhite, color.FgBlue, color.FgMagenta, color.FgBlue, color.FgYellow, color.FgGreen, color.FgRed}
	for _, c := range colors {
		c.Print("● ")
	}
	fmt.Println()
}
