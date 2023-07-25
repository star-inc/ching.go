// Nui.go
// (c) 2022 Star Inc.

package nui

type EnvKey string

func (e EnvKey) String() string {
	return string(e)
}

type IntMode int

const (
	IntModeInt IntMode = iota
	IntModeInt8
	IntModeInt16
	IntModeInt32
	IntModeInt64
)
