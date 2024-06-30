package main

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"strings"
)

func PrintBoxedText(text string) {
    lines := strings.Split(text, "\n")
    maxLength := 0
    for _, line := range lines {
        if len(line) > maxLength {
            maxLength = len(line)
        }
    }

    horizontalBorder := "+" + strings.Repeat("-", maxLength+2) + "+"

    fmt.Println(gchalk.Blue(horizontalBorder))
    for _, line := range lines {
    	// Format the line with padding
        formattedLine := fmt.Sprintf("| %-*s |", maxLength, line)
        // Split the line into parts and apply different colors
        coloredLine := gchalk.Blue(formattedLine[:1]) + gchalk.Red(formattedLine[1:maxLength+3]) + gchalk.Blue(formattedLine[maxLength+3:])
        fmt.Println(coloredLine)
    }
    fmt.Println(gchalk.Blue(horizontalBorder))
}
