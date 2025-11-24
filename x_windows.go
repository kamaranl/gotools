//go:build windows

package x

import (
	"os/exec"

	"golang.org/x/sys/windows"
)

func (a *Alert) alert() {
	var boxtype uint32
	switch a.Level {
	case Informational:
		boxtype = windows.MB_ICONINFORMATION
	case Warning:
		boxtype = windows.MB_ICONWARNING
	case Critical:
		boxtype = windows.MB_ICONERROR
	}

	_, _ = windows.MessageBox(
		0,
		windows.StringToUTF16Ptr(a.Message),
		windows.StringToUTF16Ptr(a.Title),
		windows.MB_APPLMODAL|boxtype,
	)
}

func OpenUrl(url string) error {
	return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
}
