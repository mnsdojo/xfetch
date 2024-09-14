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

// Define maximum widths for alignment
const (
	iconWidth   = 2 // Width for the icon, including some extra space
	labelWidth  = 15 // Width for the label, adjust as needed
	valueWidth  = 50 // Width for the value, adjust as needed
)

// PrintInfo prints a line of information with a label, value, and color.
func PrintInfo(label, value string, labelColor color.Color) {
	icon, ok := icons[label]
	if !ok {
		fmt.Printf("Unknown label: %s\n", label)
		return
	}

	// Print formatted line with consistent spacing
	labelColor.Printf(" %s %-"+fmt.Sprint(labelWidth)+"s → %-"+fmt.Sprint(valueWidth)+"s\n",
		icon, label, value)
}

// PrintColorDots prints a series of colored dots.
func PrintColorDots() {
	colors := []color.Color{color.FgWhite, color.FgWhite, color.FgBlue, color.FgMagenta, color.FgBlue, color.FgYellow, color.FgGreen, color.FgRed}
	for _, c := range colors {
		c.Print("● ")
	}
	fmt.Println()
}
