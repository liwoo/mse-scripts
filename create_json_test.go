package main

import "testing"

func Create_JSON_Cleaner(t *testing.T) {
	wanted := true
	csvFile := "files/csv/4115.csv"
	data, err := Clean(csvFile, "files/errors/")
	if err != nil {
		t.Fatalf("Wanted %t but got %t", wanted, false)
	}
	dailys, codeGraphs := createJson(data)
	got := len(dailys) > 0 && len(codeGraphs) > 0

	if got != wanted {
		t.Fatalf("Wanted %t but got %t", wanted, got)
	}
}
