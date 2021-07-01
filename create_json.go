package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// TODO:  next script should store following json
// dailyStats:
// 		"2020-05-19": [
//				{
//					code: "MPICO",
//					closing: 24
//					etc
//				}]

// codeGraph:
//		"MPICO": [
//			{
//				date: 44854856,
//				closing: 54.55
//			}]

type CodeGraph struct {
	Date    time.Time   `json:"date"`
	Closing float64 `json:"closing"`
}

func createJson(dailyRates []DailyCompanyRate, date string, jsonPath string) (string) {
	dailyStatsRaw, err := json.Marshal(dailyRates)
	if err != nil {
		log.Fatalf("could not convert to json")
	}

	dailyStatsJson := fmt.Sprintf("{ \"%s\": %s}", date, dailyStatsRaw)
	saveJson(dailyStatsJson, fmt.Sprintf("%s%s-daily", jsonPath, strings.ReplaceAll(date, "/", "-")))
	return dailyStatsJson
}

func saveJson(content string, fileName string) {
	jsonFile, err := os.Create(fmt.Sprintf("%s.json", fileName))

	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	_, errr := jsonFile.WriteString(content)
	if errr != nil {
		log.Fatal(err)
	}
}
