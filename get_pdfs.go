package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sirsean/go-pool"
)

//TODO: Get these from .env
var start = 3500
var end = 4117
var path = "files/pdfs/"

func main() {

	//Go Pool does the trick!!!
	p := pool.NewPool(100, 10) //values from env
	p.Start()
	for i := start; i <= end; i++ {
		p.Add(FileDownloader{
			fmt.Sprint("http://mse.co.mw/index.php?route=market/download/report&rid=", i), //Grab this from env
			fmt.Sprint(path, i, ".pdf"),
			Client,
		})
	}
	p.Close()
}

type FileDownloader struct {
	FileUrl  string
	FileName string
	Client   http.Client
}

func (u FileDownloader) Perform() {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}

	file, err := os.Create(u.FileName)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := u.Client.Get(u.FileUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	// TODO: check if size is greate than 100kb and remove the file
	//since some of them are monthly reports
	//also the ones less than 50kb (44kb) are duplicates so let's remove them too

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", u.FileName, size)
}
