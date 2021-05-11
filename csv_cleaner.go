package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type DailyCompanyRate struct {
	NO         string `json:"no"`
	HIGH       string `json:"high"`
	LOW        string `json:"low"`
	CODE       string `json:"code"`
	BUY        string `json:"buy"`
	SELL       string `json:"sell"`
	PCP        string `json:"pcp"`
	TCP        string `json:"tcp"`
	VOL        string `json:"vol"`
	DIVNET     string `json:"div_net"`
	DIVYIELD   string `json:"div_yield"`
	EEARNYIELD string `json:"earn_yield"`
	PERATIO    string `json:"pe_ratio"`
	PBVRATION  string `json:"pbv_ratio"`
	CAP        string `json:"cap"`
	PROFIT     string `json:"profit"`
	SHARES     string `json:"shares"`
}

func Clean(csvFile string, errorPath string) ([]DailyCompanyRate, []string) {
	var csvRaw []string
	var rates []DailyCompanyRate
	var errors []string

	file, err := os.Open(csvFile)

	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i int16 = 0

	for scanner.Scan() {
		//do work
		line := scanner.Text()

		if i > 26 && i < 43 {
			csvRaw = append(csvRaw, line)
		}

		i += 1
	}

	b := []byte(strings.Join(csvRaw, "\n"))

	r := csv.NewReader(bytes.NewBuffer(b))
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var rate DailyCompanyRate
	for i, word := range records {
		if len(word) == 17 {

			rate.NO = strings.TrimSpace(word[0])
			rate.HIGH = strings.TrimSpace(word[1])
			rate.LOW = strings.TrimSpace(word[2])
			rate.CODE = strings.TrimSpace(word[3])
			rate.BUY = strings.TrimSpace(word[4])
			rate.SELL = strings.TrimSpace(word[5])
			rate.PCP = strings.TrimSpace(word[6])
			rate.TCP = strings.TrimSpace(word[7])
			rate.VOL = strings.TrimSpace(word[8])
			rate.DIVNET = strings.TrimSpace(word[9])
			rate.DIVYIELD = strings.TrimSpace(word[10])
			rate.EEARNYIELD = strings.TrimSpace(word[11])
			rate.PERATIO = strings.TrimSpace(word[12])
			rate.PBVRATION = strings.TrimSpace(word[13])
			rate.CAP = strings.TrimSpace(word[14])
			rate.PROFIT = strings.TrimSpace(word[15])
			rate.SHARES = strings.TrimSpace(word[16])

			if Verify(rate) {
				rates = append(rates, rate)
			} else {
				errors = append(errors, fmt.Sprintf("line: %b values: %s", i, word))
			}
		}
	}

	if len(errors) > 0 {
		affected := fmt.Sprintf("File: %q", csvFile)
		errors = append([]string{affected}, errors...)
		logErrors(errors, errorPath)
	}

	return rates, errors
}

func logErrors(errors []string, path string) {
	f, err := os.Create(fmt.Sprintf("%s%s-error.txt", path, time.Now().Format("06-01-02 15-04-05")))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(strings.Join(errors, "\n"))
	if err2 != nil {
		log.Fatal(err)
	}
}

func Verify(rate DailyCompanyRate) bool {
	if isNotEmpty(rate.CODE) && isNotEmpty(rate.BUY) && isNotEmpty(rate.SELL) && isNotEmpty(rate.PCP) &&
		isNotEmpty(rate.TCP) && isNotEmpty(rate.VOL) && isNotEmpty(rate.DIVNET) && isNotEmpty(rate.DIVYIELD) &&
		isNotEmpty(rate.EEARNYIELD) && isNotEmpty(rate.PERATIO) && isNotEmpty(rate.PBVRATION) && isNotEmpty(rate.CAP) &&
		isNotEmpty(rate.PROFIT) && isNotEmpty(rate.SHARES) {
		return true
	}

	return false
}

func isNotEmpty(value string) bool {
	return len(value) > 0
}
