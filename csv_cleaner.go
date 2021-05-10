package main

import (
	"bufio"
	"log"
	"os"
)

func Clean(csvFile string) []string {

	file, err := os.Open(csvFile)

	if err != nil {
		log.Fatalf("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var csv []string
	headers := "No, High, Low, Code, Buy, Sell, PCP, TCP, Vol, DivNet, DivYield, EearnYield, PERatio, PBVRation, Cap, Profit, Shares"
	csv = append(csv, headers)

	var i int16 = 0

	for scanner.Scan() {
		//do work
		line := scanner.Text()

		if i > 26 && i < 41 {
			csv = append(csv, line)
		}

		i += 1
	}

	return csv
}

//TODO: Verify func that will be thoroughly tested to make sure its a csv
func Verify(csv []string) bool {
	return true
}
