package main

import "testing"

func Create_JSON_Cleaner(t *testing.T) {
	wanted := true
	csvFile := "files/csv/4115.csv"
	data := Clean(csvFile, "files/errors/")
	if len(data.errors) <= 0 {
		t.Fatalf("Wanted %t but got %t", wanted, false)
	}
	dailys, codeGraphs := createJson(data.dailyRates, data.date)
	got := len(dailys) > 0 && len(codeGraphs) > 0

	if got != wanted {
		t.Fatalf("Wanted %t but got %t", wanted, got)
	}
}
