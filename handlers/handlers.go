package handlers

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AlphaWong/html2pdf/utils"
)

func PdfHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("METHOD_NOT_ALLOW")
		http.Error(w, "METHOD_NOT_ALLOW", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, utils.MaxUploadSize)
	if err := r.ParseMultipartForm(utils.MaxUploadSize); err != nil {
		log.Println("FILE_TOO_BIG")
		log.Printf("%v \n", err)
		http.Error(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("CANNOT_FOUND_FILE")
		log.Printf("%v \n", err)
		http.Error(w, "CANNOT_FOUND_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("INVALID_FILE")
		log.Printf("%v \n", err)
		http.Error(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	filetype := http.DetectContentType(fileBytes)
	if filetype != "text/html; charset=utf-8" {
		log.Println("INVALID_FILE_TYPE")
		log.Printf("%v \n", filetype)
		http.Error(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	var fileName = utils.GetUUID()
	tmpFile, err := ioutil.TempFile(utils.UploadPath, fileName+"*.html")
	log.Printf("%v \n", tmpFile.Name())
	defer os.Remove(tmpFile.Name()) // clean up
	if err != nil {
		log.Println("CANNOT_CREATE_TMP_FILE")
		log.Printf("%v \n", err)
		http.Error(w, "CANNOT_CREATE_TMP_FILE", http.StatusInternalServerError)
		return
	}

	_, err = tmpFile.Write(fileBytes)
	if err != nil {
		log.Println("CANNOT_WRITE_TMP_FILE")
		log.Printf("%v \n", err)
		http.Error(w, "CANNOT_WRITE_TMP_FILE", http.StatusInternalServerError)
		return
	}

	tmpFile.Sync()

	if err := tmpFile.Close(); err != nil {
		log.Println("CANNOT_CLOSE_TMP_FILE")
		log.Printf("%v \n", err)
		http.Error(w, "CANNOT_CLOSE_TMP_FILE", http.StatusInternalServerError)
	}

	var pdfFileFullPath = utils.PdfPath + fileName + ".pdf"
	var cp = &utils.ConverterParam{
		InFilePath:  tmpFile.Name(),
		OutFilePath: pdfFileFullPath,
	}
	var c = utils.NewConverter(cp)
	if err := c.ConvertHtml2Pdf(); err != nil {
		log.Println("CANNOT_CONVERT_PDF")
		log.Printf("%v \n", err)
		http.Error(w, "CANNOT_CONVERT_PDF", http.StatusInternalServerError)
		return
	}

	// create header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/pdf")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Session-Id", fileName)

	b, err := ioutil.ReadFile(pdfFileFullPath)
	gz := gzip.NewWriter(w)
	gz.Write(b)
	defer gz.Close()
	os.Remove(pdfFileFullPath)
}
