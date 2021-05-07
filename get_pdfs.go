package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pdftables/go-pdftables-api/pkg/client"
	"github.com/sirsean/go-pool"
	"github.com/spf13/viper"
)

var CONFIG Configuration
var clientCSV client.Client

func initConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config")
	}

	err := viper.Unmarshal(&CONFIG)

	if err != nil {
		log.Fatal()
	}

	clientCSV = client.Client{
		APIKey:     CONFIG.PDFTABLES_API_KEY,
		HTTPClient: http.DefaultClient,
	}
}

func folderCheck(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
	}
}

func main() {
	initConfig()
	folderCheck(CONFIG.RAW_PDF_PATH)
	folderCheck(CONFIG.RAW_CSV_PATH)
	//Go Pool does the trick!!!
	p := pool.NewPool(CONFIG.QUEUE_SIZE, CONFIG.WORKER_NUM)
	p.Start()
	for i := CONFIG.START; i <= CONFIG.END; i++ {
		p.Add(FileDownloader{
			fmt.Sprint(CONFIG.MSE_URL, i),
			fmt.Sprint(CONFIG.RAW_PDF_PATH, i, ".pdf"),
			fmt.Sprint(CONFIG.RAW_CSV_PATH, i, ".csv"),
			Client,
		})
	}
	p.Close()
}

type FileDownloader struct {
	FileUrl     string
	FileName    string
	FileNameCSV string
	Client      http.Client
}

func (u FileDownloader) Perform() {
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
		} else {
			func() {
				file, err := os.Open(u.FileName)

				if err != nil {
					log.Fatal(err)
				}

				defer file.Close()
				csvFile, err := os.Create(u.FileNameCSV)
				if err != nil {
					log.Fatal(err)
				}

				converted, err := clientCSV.Do(file, client.FormatCSV)
				if err != nil {
					log.Fatal(err)
				}

				_, err = io.Copy(csvFile, converted)
				if err != nil {
					log.Fatal(err)
				}

				defer csvFile.Close()
			}()
		}
	}()

	fmt.Printf("Downloaded a file %s with size %d\n", u.FileName, size)
}
