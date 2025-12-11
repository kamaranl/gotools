package x

import (
	"strings"
)

var alerts = NewTypeSafeMap[*Alert]()

type AlertLevel int

const (
	Informational AlertLevel = iota
	Warning
	Critical
)

type Alert struct {
	Title   string
	Message string
	Level   AlertLevel
	OnExit  func(ok bool)

	ok    bool
	label string
}

func NewAlert(title, message string, level AlertLevel) *Alert {
	label := strings.ToLower(strings.ReplaceAll(title, " ", ""))

	return &Alert{
		Title:   title,
		Message: message,
		Level:   level,
		label:   "alert://" + label,
	}
}

func (a *Alert) Show() {
	if alert, ok := alerts.Get(a.label); ok && alert != nil {
		return
	}

	alerts.Set(a.label, a)
	go func() {
		a.alert()
		if a.OnExit != nil {
			a.OnExit(a.ok)
		}
		alerts.Delete(a.label)
	}()
}
