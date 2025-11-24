package x

import (
	"strings"

	xstate "github.com/kamaranl/x/state"
)

var state = xstate.NewState()

type Alert struct {
	Title   string
	Message string
	Level   AlertLevel
	label   string
}

func NewAlert(title, message string, level AlertLevel) *Alert {
	label := strings.ToLower(strings.ReplaceAll(title, " ", ""))

	return &Alert{
		Title:   title,
		Message: message,
		Level:   level,
		label:   "alert&" + label,
	}
}

func (a *Alert) Show() {
	if active, ok := xstate.Get[bool](state, a.label); ok && active {
		return
	}
	state.Set(a.label, true)

	go func() {
		a.alert()
		state.Set(a.label, false)
	}()
}

type AlertLevel int

const (
	Informational AlertLevel = iota
	Warning
	Critical
)
