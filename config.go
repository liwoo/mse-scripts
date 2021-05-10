package main
type Configuration struct {
	MSE_URL string
	PDFTABLES_API_KEY string
	START int
	END int
	RAW_PDF_PATH string
	RAW_CSV_PATH string
	QUEUE_SIZE int
	WORKER_NUM int
}
