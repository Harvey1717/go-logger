package log

import (
	"fmt"
	"time"
	"github.com/fatih/color"
)

var formattedTime string = time.Now().Format("15:04:05.00000")
var yellow = color.New(color.FgYellow).SprintFunc()

func L(msg string) {
	fmt.Printf("[%s] - %s\n", yellow(formattedTime), yellow(msg))
}

func S(msg string) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("[%s] - %s\n", yellow(formattedTime), green(msg))
}

func E(msg string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("[%s] - %s\n", yellow(formattedTime), red(msg))
}