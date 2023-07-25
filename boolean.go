// nui.go
// (c) 2023 Star Inc.

package nui

func IsYes(key EnvKey) bool {
	yesString := String(key)

	return yesString == "yes"
}
