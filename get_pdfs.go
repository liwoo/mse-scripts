package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var start = 3500
var end = 4117
var path = "files/pdfs/"

func main() {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	for i := start; i <= end; i++ {
		getPdf(i, &client)
	}
}

func getPdf(pos int, client *http.Client) {
	fileLocation := fmt.Sprint("http://mse.co.mw/index.php?route=market/download/report&rid=", pos)

	fileName := fmt.Sprint(path, pos, ".pdf")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(fileLocation)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}
