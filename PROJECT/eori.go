package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type EORI struct {
	Value string
}

func (e *EORI) Process(line string) (string, uint64) {
	r := regexp.MustCompile(`EORI\ +\#0X([ABCDEF\d]{0,2})$`)
	if results := r.FindStringSubmatch(strings.ToUpper(line)); len(results) > 0 {
		e.Value = results[1]
		value := e.eori()
		return value, calculateChecksum(value)
	}

	return "", 0
}

func (e *EORI) eori() string {
	number, _ := strconv.ParseUint(e.Value, 16, 64)
	binary := "0000101000111100" + fmt.Sprintf("%08s", strconv.FormatInt(int64(number), 2))

	ui, _ := strconv.ParseUint(binary, 2, 64)
	return fmt.Sprintf("%x", ui)
}
