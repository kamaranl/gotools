//go:build windows

package dialog

import (
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

	ret, _ := windows.MessageBox(
		0,
		windows.StringToUTF16Ptr(a.Message),
		windows.StringToUTF16Ptr(a.Title),
		windows.MB_APPLMODAL|windows.MB_OKCANCEL|boxtype,
	)

	if ret == 1 {
		a.ok = true
	}
}
