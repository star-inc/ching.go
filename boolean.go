// Nui.go
// (c) 2022 Star Inc.

package nui

func IsYes(key EnvKey) bool {
	yesString := Get(key)

	return yesString == "yes"
}
