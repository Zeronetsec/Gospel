// Gospel Project

package invinput

import (
    "fmt"
    "gospel/utils/color"
)

func Invalid() {
    fmt.Printf(
        "%s[!] %sInvalid input!\n",
        color.R, color.N,
    )

    fmt.Printf(
        "%s[!] %sTry: %sgospel --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Zeronetsec