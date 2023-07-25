// nui.go
// (c) 2023 Star Inc.

package nui

import (
	"log"
	"strconv"
)

func GetInt(key EnvKey, mode IntMode) any {
	intString := Get(key)

	i, err := strconv.ParseInt(intString, 10, 64)
	if err != nil {
		log.Fatalf("nui invalid int value: %s", intString)
	}

	switch mode {
	case IntModeInt:
		return int(i)
	case IntModeInt8:
		return int8(i)
	case IntModeInt16:
		return int16(i)
	case IntModeInt32:
		return int32(i)
	case IntModeInt64:
		return int64(i)
	default:
		log.Fatalf("nui invalid int mode: %d", mode)
		return nil
	}
}
