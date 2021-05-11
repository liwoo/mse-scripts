package main

import "testing"

func Test_CSV_Cleaner(t *testing.T) {
	csvFile := "files/csv/4115.csv"
	_, err := Clean(csvFile, "files/errors/")
	got := len(err) > 0
	wanted := false

	if got != wanted {
		t.Fatalf("Wanted %t but got %t", wanted, got)
	}
}

//TODO Verify Edge Cases
