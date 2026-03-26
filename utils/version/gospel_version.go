// Gospel Project

package version

import (
    "fmt"
    "gospel/utils/color"
)

const (
    tools = "Gospel"
    version = "v1.0"
    homepage = "https://github.com/Zeronetsec/Gospel"
)

func GospelVersion() {
    fmt.Printf(
        "%sProject: %s%s%s\n",
        color.N, color.GG, tools, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}