package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type BKPT struct {
	Value string
}

func (e *BKPT) Process(line string) (string, uint64) {
	r := regexp.MustCompile(`BKPT\ +\#0X([0-7])$`)
	if results := r.FindStringSubmatch(strings.ToUpper(line)); len(results) > 0 {
		e.Value = results[1]
		value := e.bkpt()
		return value, calculateChecksum(value)
	}

	return "", 0
}

func (e *BKPT) bkpt() string {
	number, _ := strconv.ParseUint(e.Value, 16, 64)
	binary := "0100100001001" + fmt.Sprintf("%03s", strconv.FormatInt(int64(number), 2))

	ui, _ := strconv.ParseUint(binary, 2, 64)
	return fmt.Sprintf("%x", ui)
}
