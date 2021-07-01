package main

import (
	"log"

	"github.com/spf13/viper")


type Configuration struct {
	MSE_URL           string
	PDFTABLES_API_KEY string
	DOWLOADSTART             int
	DOWNLOADEND               int
	RAW_PDF_PATH      string
	RAW_CSV_PATH      string
	ERROR_FILE_PATH   string
	CLEANED_CSV_PATH  string
	CLEANED_JSON_PATH string
	QUEUE_SIZE        int
	WORKER_NUM        int
}

func initConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config")
	}

	err := viper.Unmarshal(&CONFIG)

	if err != nil {
		log.Fatal()
	}

	
}