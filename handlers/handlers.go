package handlers

import (
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/AlphaWong/html2pdf/utils"
	"github.com/lalamove-go/logs"
)

func PdfHandler(w http.ResponseWriter, r *http.Request) {
	// Generate UUID v4 for the file
	// It is the session id also
	var fileName = utils.GetUUID()

	// Add session id to header
	w.Header().Set("Session-Id", fileName)

	if r.Method != http.MethodPost {
		logs.Logger().Error(utils.ErrorMethodNotAllow, zap.String(utils.SessionID, fileName))
		http.Error(w, utils.ErrorMethodNotAllow, http.StatusMethodNotAllowed)
		return
	}

	// Check file size
	r.Body = http.MaxBytesReader(w, r.Body, utils.MaxUploadSize)
	if err := r.ParseMultipartForm(utils.MaxUploadSize); err != nil {
		logs.Logger().Error(utils.ErrorFileTooBig, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorFileTooBig, http.StatusBadRequest)
		return
	}

	// Check file existence
	file, _, err := r.FormFile("file")
	if err != nil {
		logs.Logger().Error(utils.ErrorFileNotFound, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorFileNotFound, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Log form values
	logs.Logger().Info("form data", zap.String(utils.SessionID, fileName), zap.Any("value", r.MultipartForm.Value))

	// Read file to byte
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logs.Logger().Error(utils.ErrorInvalidFile, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorInvalidFile, http.StatusBadRequest)
		return
	}

	// Check file type
	filetype := http.DetectContentType(fileBytes)
	if filetype != "text/html; charset=utf-8" {
		logs.Logger().Error(utils.ErrorInvalidFileType, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorInvalidFileType, http.StatusBadRequest)
		return
	}

	// Create tmp file
	tmpFile, err := ioutil.TempFile(utils.UploadPath, fileName+"*.html")
	log.Printf("%v \n", tmpFile.Name())
	if err != nil {
		logs.Logger().Error(utils.ErrorCannotCreateTmpFile, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorCannotCreateTmpFile, http.StatusInternalServerError)
		return
	}

	// Clean up tmp file
	defer os.Remove(tmpFile.Name())

	// Write byte to tmp file
	_, err = tmpFile.Write(fileBytes)
	if err != nil {
		logs.Logger().Error(utils.ErrorCannotWriteTmpFile, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorCannotWriteTmpFile, http.StatusInternalServerError)
		return
	}

	// Force write to disk
	// reference : https://www.joeshaw.org/dont-defer-close-on-writable-files/
	// Credit Alan TANG
	tmpFile.Sync()

	// Close the tmp file after writing
	if err := tmpFile.Close(); err != nil {
		logs.Logger().Error(utils.ErrorCannotCloseTmpFile, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorCannotCloseTmpFile, http.StatusInternalServerError)
	}

	// Generate pdf file path
	var pdfFileFullPath = utils.PdfPath + fileName + ".pdf"
	var cp = &utils.ConverterParam{
		InFilePath:  tmpFile.Name(),
		OutFilePath: pdfFileFullPath,
		Options:     utils.ParseFormValues(r.MultipartForm.Value),
	}

	// Convert html to pdf
	var c = utils.NewConverter(cp)
	if err := c.ConvertHtml2Pdf(); err != nil {
		logs.Logger().Error(utils.ErrorCannotConvertPDF, zap.String(utils.SessionID, fileName), zap.Error(err))
		http.Error(w, utils.ErrorCannotConvertPDF, http.StatusInternalServerError)
		return
	}

	// Create header
	w.Header().Add("Accept-Charset", "utf-8")
	w.Header().Add("Content-Type", "application/pdf")
	w.Header().Set("Content-Encoding", "gzip")

	// Convert pdf file in gzip format
	b, _ := ioutil.ReadFile(pdfFileFullPath)
	gz := gzip.NewWriter(w)
	gz.Write(b)
	defer gz.Close()

	// Clear up pdf file
	os.Remove(pdfFileFullPath)
}
