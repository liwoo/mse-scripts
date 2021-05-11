package main

import "testing"

func Test_CSV_Cleaner(t *testing.T) {
	csvFile := "files/csv/4115.csv"
	got := Verify(Clean(csvFile))
	wanted := true

	if got != wanted {
		t.Fatalf("Wanted %t but got %t", wanted, got)
	}
}

//TODO Verify Edge Cases
