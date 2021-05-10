package main

import "testing"

func Test_CSV_Cleaner(t *testing.T) {
	csvFile := "files/csvs/3500.csv"
	got := Verify(Clean(csvFile))
	wanted := true

	if got != wanted {
		t.Fatalf("Wanted %t but got %t", wanted, got)
	}
}

//TODO Verify Edge Cases
