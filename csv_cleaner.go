package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type DailyCompanyRate struct {
	NO string `json:"no"`
	HIGH string `json:"high"`
	LOW string `json:"low"`
	CODE string `json:"code"`
	BUY string `json:"buy"`
	SELL string `json:"sell"`
	PCP string `json:"pcp"`
	TCP string `json:"tcp"`
	VOL string `json:"vol"`
	DIVNET string `json:"div_net"`
	DIVYIELD string `json:"div_yield"`
	EEARNYIELD string `json:"earn_yield"`
	PERATIO string `json:"pe_ratio"`
	PBVRATION string `json:"pbv_ratio"`
	CAP string `json:"cap"`
	PROFIT string `json:"profit"`
	SHARES string `json:"shares"`
} 

func Clean(csvFile string) []DailyCompanyRate {
	var csvRaw []string
	var rates []DailyCompanyRate

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

		if i > 26 && i < 41 {
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
	for _, word := range records {
		if len(word) == 17 {

			rate.NO = word[0]
			rate.HIGH = word[1]
			rate.LOW = word[2]
			rate.CODE = word[3]
			rate.BUY = word[4]
			rate.SELL = word[5]
			rate.PCP = word[6]
			rate.TCP = word[7]
			rate.VOL = word[8]
			rate.DIVNET = word[9]
			rate.DIVYIELD = word[10]
			rate.EEARNYIELD = word[11]
			rate.PERATIO = word[12]
			rate.PBVRATION = word[13]
			rate.CAP = word[14]
			rate.PROFIT = word[15]
			rate.SHARES = word[16]

			rates = append(rates, rate)
		}
	}
	return rates
}

func Verify(dailyRates []DailyCompanyRate) bool {
	for _, rate := range dailyRates {
		if len(rate.CODE) <= 0 {
			return false
		}
	}
	return true
}
