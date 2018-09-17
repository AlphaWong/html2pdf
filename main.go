package main

import (
	"log"
	"net/http"

	"github.com/AlphaWong/html2pdf/boot"
	"github.com/AlphaWong/html2pdf/handlers"
	"github.com/AlphaWong/html2pdf/utils"
	"github.com/lalamove-go/logs"
	"go.uber.org/zap"
)

func main() {
	boot.Init()

	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/convert", handlers.PdfHandler)

	logs.Logger().Info("Server on at", zap.String("PORT", utils.Port))
	log.Fatal(http.ListenAndServe(utils.Port, nil))
}
