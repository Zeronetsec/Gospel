// Gospel Project

package main

import (
    "os"
    "strings"
    "gospel/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.GospelConsole(input)
}

// Copyright (c) 2026 Zeronetsec