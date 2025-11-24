//go:build darwin

package x

import (
	"fmt"
	"os/exec"
)

func (a *Alert) alert() {
	var level string
	switch a.Level {
	case Informational:
		level = "informational"
	case Warning:
		level = "warning"
	case Critical:
		level = "critical"
	}

	format := "display alert %q message %q as %s"
	script := fmt.Sprintf(format, a.Title, a.Message, level)
	_ = exec.Command("osascript", "-e", script).Start()
}

func OpenUrl(url string) error { return exec.Command("open", url).Start() }
