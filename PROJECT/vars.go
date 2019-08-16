package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VARS struct {
	Value string
}

func (v *VARS) Process(line string) (string, uint64) {
	r := regexp.MustCompile(`VAR\d\ +(BYTE|WORD|LONG|DDATA\.?[BWL]?)\ +\$([ABCDEF\d]+)$`)
	if results := r.FindStringSubmatch(strings.ToUpper(line)); len(results) > 0 {
		v.Value = results[2]
		switch results[1] {
		case "BYTE":
			value := v.simple(2)
			return value, calculateChecksum(value)
		case "WORD":
			value := v.simple(4)
			return value, calculateChecksum(value)
		case "LONG":
			value := v.simple(8)
			return value, calculateChecksum(value)
		case "DDATA.B":
			value := v.data("00")
			return value, calculateChecksum(value)
		case "DDATA.L":
			value := v.data("00000000")
			return value, calculateChecksum(value)
		default:
			value := v.data("0000")
			return value, calculateChecksum(value)
		}
	}

	return "", 0
}

func (v *VARS) simple(size int) string {
	var extra int
	zero := (len(v.Value) % size)
	if zero != 0 {
		extra = size - zero
	}
	return fmt.Sprintf(fmt.Sprintf("%s0%ds", "%", len(v.Value)+extra), v.Value)
}

func (v *VARS) data(size string) string {
	var index uint64
	var result string
	ui, _ := strconv.ParseUint(v.Value, 16, 64)
	for index = 0; index < ui; index++ {
		result += size
	}

	return result
}
