package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AlphaWong/html2pdf/handlers"
	"github.com/AlphaWong/html2pdf/utils"
)

func main() {
	boot()

	http.HandleFunc("/convert", handlers.PdfHandler)

	log.Println("Server on at", utils.Port)
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}

func boot() {
	var maxSize, err = strconv.ParseInt(os.Getenv("MAX_SIZE"), 10, 64)
	if err != nil {
		log.Fatal("Invalid MAX_SIZE env")
	}
	utils.MaxUploadSize = maxSize * 1024 * 1024 // MB as unit
}
