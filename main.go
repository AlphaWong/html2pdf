package main

import (
	"log"
	"net/http"

	"github.com/AlphaWong/html2pdf/boot"
	"github.com/AlphaWong/html2pdf/handlers"
	"github.com/AlphaWong/html2pdf/utils"
)

func main() {
	boot.Init()

	http.HandleFunc("/convert", handlers.PdfHandler)

	log.Println("Server on at", utils.Port)
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}
