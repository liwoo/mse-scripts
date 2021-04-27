package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//TODO: Use Range from 3500 to 4117 to create these dynamically
	//after being defined up top
	fileLocation := "http://mse.co.mw/index.php?route=market/download/report&rid=3500"
	//file location should be in files/pdfs/... and then files/csv/...
	fileName := "3500.pdf"

	//Keep these in a function
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	//Reuse the client
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fileLocation)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	//To act like our log
	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}
