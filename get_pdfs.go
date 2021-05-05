package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sirsean/go-pool"
	"github.com/spf13/viper"
)

var CONFIG Configuration

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

func main() {
	initConfig()
	//Go Pool does the trick!!!
	p := pool.NewPool(CONFIG.QUEUE_SIZE, CONFIG.WORKER_NUM)
	p.Start()
	for i := CONFIG.START; i <= CONFIG.END; i++ {
		p.Add(FileDownloader{
			fmt.Sprint(CONFIG.MSE_URL, i),
			fmt.Sprint(CONFIG.RAW_PDF_PATH, i, ".pdf"),
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
	if _, err := os.Stat(CONFIG.RAW_PDF_PATH); os.IsNotExist(err) {
		os.MkdirAll(CONFIG.RAW_PDF_PATH, 0700)
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

	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}

		if size < 49000 || size > 80000 {
			e := os.Remove(u.FileName)
			if e != nil {
				log.Fatal(e)
			}
		}

		//do csv conversion
		
		//close
	}()

	fmt.Printf("Downloaded a file %s with size %d\n", u.FileName, size)
}
