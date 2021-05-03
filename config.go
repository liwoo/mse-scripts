package main


type Configuration struct {
	MSE_URL string
	START int
	END int
	RAW_PDF_PATH string
	QUEUE_SIZE int
	WORKER_NUM int
}
