// Gospel Project

package console

import (
    "time"
    "gospel/utils/invinput"
    "gospel/utils/uwu"
    "gospel/utils/version"
    "gospel/utils/helper"
    "gospel/utils/cursor"
    "gospel/module/sysinfo"
    "gospel/module/procinfo"
    "gospel/module/checkroot"
    "gospel/module/dumpstring"
    "gospel/module/decode"
    "gospel/module/misconfind"
)

type Command interface {
    Execute(args []string)
}

type Sysinfo struct{}
func (c Sysinfo) Execute(args []string) {
    cursor.Hide()
    sysinfo.MachineInfo()
    cursor.Visible()
}

type Procinfo struct{}
func (c Procinfo) Execute(args []string) {
    cursor.Hide()
    procinfo.ShowProcInfo()
    cursor.Visible()
}

type Checkroot struct{}
func (c Checkroot) Execute(args []string) {
    cursor.Hide()
    checkroot.GetCheck()
    cursor.Visible()
}

type Dumpstring struct{}
func (c Dumpstring) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    dumpstring.Analyzer(args[2])
    cursor.Visible()
}

type Decode struct{}
func (c Decode) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    decode.ExecBrute(args[2])
    cursor.Visible()
}

type Misconfind struct{}
func (c Misconfind) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgs()
        return
    }

    cursor.Hide()
    misconfind.Finder(args[2:])
    cursor.Visible()
}

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Nyanners(5 * time.Second)
    cursor.Visible()
}

type Version struct{}
func (c Version) Execute(args []string) {
    cursor.Hide()
    version.GospelVersion()
    cursor.Visible()
}

type Helper struct{}
func (c Helper) Execute(args []string) {
    cursor.Hide()
    helper.GospelHelper()
    cursor.Visible()
}

// Copyright (c) 2026 Zeronetsec