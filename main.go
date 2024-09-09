package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

var icons = map[string]string{
	"OS":      "🍎",
	"Machine": "□",
	"Kernel":  "⚙",
	"Uptime":  "◷",
	"Shell":   "⌘",
	"CPU":     "⏣",
	"Memory":  "▤",
}

func main() {
	hostname, _ := os.Hostname()
	color.Cyan.Printf("\n%s@%s\n\n", os.Getenv("USER"), hostname)
	fmt.Println("┌─────────────────── System Information ───────────────────┐")
	
	

}


