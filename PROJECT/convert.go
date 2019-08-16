package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getBinnaryRegister(register string) string {
	value := strings.Split(strings.ToUpper(register), "")
	if value[0] == "A" || value[0] == "D" {
		number, _ := strconv.Atoi(value[1])
		return fmt.Sprintf("%03s", strconv.FormatInt(int64(number), 2))
	}

	return ""
}

func getBinnaryRegisterWithMode(register string) (string, string) {
	value := strings.Split(strings.ToUpper(register), "")
	if value[0] == "D" {
		return "000", getBinnaryRegister(register)
	} else if value[0] == "(" {
		register = strings.Replace(register, "(", "", -1)
		register = strings.Replace(register, ")", "", -1)
		if strings.Contains(register, "+") {
			register = strings.Replace(register, "+", "", -1)
			return "011", getBinnaryRegister(register)
		}
		return "010", getBinnaryRegister(register)
	} else if value[0] == "-" {
		register = strings.Replace(register, "-", "", -1)
		register = strings.Replace(register, "(", "", -1)
		register = strings.Replace(register, ")", "", -1)
		return "100", getBinnaryRegister(register)
	}

	return "000", "000"
}
