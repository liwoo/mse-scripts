package main

import (
	"flag"
	"log"
)

var CONFIG Configuration

func main() {
	initConfig()
	var script string
	flag.StringVar(&script, "script", "download", "The script to run, the options are download or clean. Default: download.")

	flag.Parse()

	switch script {
	case "download":
		GetPDFS()
	case "clean":
		CleanDownloadedCSV()
	default:
		log.Fatal("Could not find specified script: ", script)
	}
}
