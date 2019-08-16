package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DIVU struct {
	Source      string
	Destination string
}

func (a *DIVU) Process(line string) (string, uint64) {
	r := regexp.MustCompile(`DIVUL?\.?(W|L)?\ +((\([A]\d\)\+?)|(\-\([A]\d\))|(D\d)),\ +((\([A]\d\)\+?)|(\-\([A]\d\))|(D\d))$`)
	if results := r.FindStringSubmatch(strings.ToUpper(line)); len(results) > 0 {
		a.Source = results[2]
		a.Destination = results[6]
		switch results[1] {
		case "L":
			l := a.long()
			return l, calculateChecksum(l)
		default:
			w := a.word()
			return w, calculateChecksum(w)
		}
	}

	return "", 0
}

func (a *DIVU) word() string {
	_, source := getBinnaryRegisterWithMode(a.Source)
	mode, destination := getBinnaryRegisterWithMode(a.Destination)
	binary := "1000" + source + "011" + mode + destination

	ui, _ := strconv.ParseUint(binary, 2, 64)
	return fmt.Sprintf("%x", ui)
}

func (a *DIVU) long() string {
	mode, destination := getBinnaryRegisterWithMode(a.Destination)
	binary := "0100110001" + mode + destination

	ui, _ := strconv.ParseUint(binary, 2, 64)
	return fmt.Sprintf("%x", ui)
}
