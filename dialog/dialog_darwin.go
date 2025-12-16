//go:build darwin

package dialog

import (
	"fmt"
	"os/exec"
)

func (a *Alert) alert() {
	var level string
	switch a.Level {
	case Informational:
		level = "note"
	case Warning:
		level = "caution"
	case Critical:
		level = "stop"
	}

	format := "display dialog %q with title %q with icon %s"
	script := fmt.Sprintf(format, a.Message, a.Title, level)
	cmd := exec.Command("osascript", "-e", script)
	if _, err := cmd.Output(); err != nil {
		return
	}
	a.ok = cmd.ProcessState.Success()
}
