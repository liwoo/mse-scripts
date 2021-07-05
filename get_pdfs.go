package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pdftables/go-pdftables-api/pkg/client"
	"github.com/sirsean/go-pool"
)

var clientCSV client.Client

func folderCheck(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0700)
	}
}

func GetPDFS() {
	clientCSV = client.Client{
		APIKey:     CONFIG.PDFTABLES_API_KEY,
		HTTPClient: http.DefaultClient,
	}
	folderCheck(CONFIG.RAW_PDF_PATH)
	folderCheck(CONFIG.RAW_CSV_PATH)
	folderCheck(CONFIG.ERROR_FILE_PATH)
	folderCheck(CONFIG.CLEANED_CSV_PATH)
	folderCheck(CONFIG.CLEANED_JSON_PATH)

	p := pool.NewPool(CONFIG.QUEUE_SIZE, CONFIG.WORKER_NUM)
	p.Start()

	for i := CONFIG.PDF_START_NO; i <= CONFIG.PDF_END_NO; i++ {
		p.Add(MSEFileDownloader{
			fmt.Sprint(CONFIG.MSE_URL, i),
			fmt.Sprint(CONFIG.RAW_PDF_PATH, i, ".pdf"),
			fmt.Sprint(CONFIG.RAW_CSV_PATH, i, ".csv"),
			Client,
		})
	}

	p.Close()
}

type MSEFileDownloader struct {
	FileUrl     string
	FileName    string
	FileNameCSV string
	Client      http.Client
}

func (u MSEFileDownloader) Perform() {
	notStandardDailyReport := func(size int64) bool {
		return size < 49000 || size > 80000
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

	if err != nil {
		log.Fatal(err)
	}

	defer func() {

		cerr := file.Close()
		if err == nil {
			err = cerr
		}

		//TODO: Should read if not standard daily report instead
		if notStandardDailyReport(size) {
			e := os.Remove(u.FileName)
			if e != nil {
				log.Fatal(e)
			}
		}

		if !notStandardDailyReport(size) {
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

				defer csvFile.Close()

				converted, err := clientCSV.Do(file, client.FormatCSV)

				if err != nil {
					log.Fatal(err)
				}

				_, err = io.Copy(csvFile, converted)

				if err != nil {
					log.Fatal(err)
				}

				// TODO: Clean then verify, if true save in another folder
				//if not, write in a txt which ones couldn't be cleaned so
				//we do so manually

			}()
		}
	}()

	fmt.Printf("Downloaded a file %s with size %d\n", u.FileName, size)
}

// TODO:  next script should store following json
// dailyStats:
// 		"2020-05-19": [
//				{
//					code: "MPICO",
//					closing: 24
//					etc
//				}]

// codeGraph:
//		"MPICO": [
//			{
//				date: 44854856,
//				closing: 54.55
//			}]
