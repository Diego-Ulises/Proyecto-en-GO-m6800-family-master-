package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	var checksum uint64

	lines, name := readFile(os.Args[1])
	fmt.Println(strings.ToUpper(name))
	for _, line := range lines {
		processLine, _ := Process(line)
		numPars := uint64((len(processLine) / 2) + 3)
		fin := calculateChecksum(fmt.Sprintf("%02x", numPars) + fmt.Sprintf("%04x", checksum) + processLine)
		result += fmt.Sprintf("S1 %02x %04x %s %02x\n", numPars, checksum, processLine, fin)

		checksum += numPars
	}

	fmt.Println(strings.ToUpper(result))
	fmt.Println(strings.ToUpper(porfin(len(lines))))
}

func porfin(numLines int) string {
	return fmt.Sprintf("S903 %04x %02x", numLines, calculateChecksum("03"+fmt.Sprintf("%x", numLines)))
}

func Process(line string) (string, uint64) {
	var vars VARS
	var divu DIVU
	var eori EORI
	var bkpt BKPT

	if l, c := vars.Process(line); len(l) > 0 {
		return l, c
	} else if l, c := divu.Process(line); len(l) > 0 {
		return l, c
	} else if l, c := eori.Process(line); len(l) > 0 {
		return l, c
	} else if l, c := bkpt.Process(line); len(l) > 0 {
		return l, c
	} else {
		fmt.Println("HAY UN ERROR")
		fmt.Println(line)
		os.Exit(1)
	}

	return "", 0
}

func calculateChecksum(line string) uint64 {
	var values []string
	var value uint64
	for i := 0; i < len(line)-1; i = i + 2 {
		values = append(values, string(line[i])+string(line[i+1]))
	}
	for _, v := range values {
		ui, _ := strconv.ParseUint(v, 16, 64)
		value += ui
	}

	if value > 255 {
		lhex := fmt.Sprintf("%x", value)
		ui, _ := strconv.ParseUint(string(lhex[len(lhex)-2])+string(lhex[len(lhex)-1]), 16, 64)
		value = 255 - ui
	} else {
		value = 255 - value
	}

	var ui uint64
	lhex := fmt.Sprintf("%x", value)
	if len(lhex) > 1 {
		ui, _ = strconv.ParseUint(string(lhex[len(lhex)-2])+string(lhex[len(lhex)-1]), 16, 64)
	} else {
		ui, _ = strconv.ParseUint(string(lhex[len(lhex)-1]), 16, 64)
	}
	return ui
}

func readFile(nameFile string) ([]string, string) {
	var lines []string

	file, err := os.Open(nameFile)
	checkError(err)
	defer file.Close()

	var l string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l = strings.TrimSpace(scanner.Text())

		if !strings.Contains(l, ".DATA") &&
			!strings.Contains(l, ".CODE") &&
			!strings.Contains(l, "END") &&
			len(l) > 0 {
			lines = append(lines, l)
		}
	}

	err = scanner.Err()
	checkError(err)

	var nameCode string
	var pars, check int
	for _, v := range nameFile {
		nameCode += fmt.Sprintf("%02x", int(v))
		check += int(v)
		pars++
	}

	pars += 2
	check += pars
	l = fmt.Sprintf("%x", check)
	ui, _ := strconv.ParseUint(string(l[len(l)-2])+string(l[len(l)-1]), 16, 64)

	value := 255 - ui

	l = fmt.Sprintf("%x", value)
	ui, _ = strconv.ParseUint(string(l[len(l)-2])+string(l[len(l)-1]), 16, 64)

	return lines, fmt.Sprintf("S0%02x0000%s%x", pars, nameCode, ui)
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
