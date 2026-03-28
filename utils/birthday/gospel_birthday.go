// Gospel Project

package birthday

import (
    "fmt"
    "time"
    "gospel/utils/color"
)

func GospelBirthDay() {
    birthDate := "03-27"
    now := time.Now().Format("01-02")
    if now == birthDate {
        fmt.Printf(
            "%s› %sHappy birthday for %sgospel %s🎉\n",
            color.R, color.N, color.GG, color.N,
        )
        fmt.Println()
    }
}

// Copyright (c) 2026 Zeronetsec