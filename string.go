// nui.go
// (c) 2023 Star Inc.

package nui

import (
	"log"
	"os"
	"strings"
)

func Get(key EnvKey) string {
	value, isKeyExists := os.LookupEnv(key.String())
	if !isKeyExists {
		log.Fatalf("nui lookup failed: %s", key)
	}
	return value
}

func GetSplited(key EnvKey) []string {
	joinedString := Get(key)
	data := strings.Split(joinedString, ",")
	for i, j := range data {
		data[i] = strings.TrimSpace(j)
	}
	return data
}
