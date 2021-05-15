package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
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
	Date    int64   `json:"date"`
	Closing float64 `json:"closing"`
}

func createJson(dailyRates []DailyCompanyRate, date string, jsonPath string) (string, string) {
	dailyStatsRaw, err := json.Marshal(dailyRates)
	var codeGraphsRaw []string
	if err != nil {
		log.Fatalf("could not convert to json")
	}

	dailyStatsJson := fmt.Sprintf("{ \"%s\": %s}", date, dailyStatsRaw)

	for _, rate := range dailyRates {
		graph := &CodeGraph{
			Date:    45554,
			Closing: rate.SHARES,
		}
		graphJson, err := json.Marshal(graph)
		if err != nil {
			log.Fatalf("could not convert to json")
		}
		value := fmt.Sprintf("{\"%s\" : [%s]}", rate.CODE, string(graphJson))
		codeGraphsRaw = append(codeGraphsRaw, value)
	}
	codeGraphsJson := fmt.Sprintf("[%s]", strings.Join(codeGraphsRaw, ","))
	saveJson(dailyStatsJson, fmt.Sprintf("%s%s-daily", jsonPath, strings.ReplaceAll(date, "/", "-")))
	saveJson(codeGraphsJson, fmt.Sprintf("%s%s-graph", jsonPath, strings.ReplaceAll(date, "/", "-")))
	return dailyStatsJson, codeGraphsJson
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
