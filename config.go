package main
type Configuration struct {
	MSE_URL string
	PDFTABLES_API_KEY string
	START int
	END int
	RAW_PDF_PATH string
	RAW_CSV_PATH string
	ERROR_FILE_PATH string
	CLEANED_CSV_PATH string
	CLEANED_JSON_PATH string
	QUEUE_SIZE int
	WORKER_NUM int
}
