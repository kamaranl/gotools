package test

import (
	"errors"
	"time"
)

const (
	ErrUnexpectedF = "unexpected error: %v"
	ErrWantFGotF   = "wanted '%v', got '%v'"
)

const (
	TestsDisabled = " tests are disabled"
)

var Err_ = errors.New("")

type Scene struct {
	Input   any
	Output  any
	Passing bool
}

func Countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		println(i)
		time.Sleep(1 * time.Second)
	}
}
