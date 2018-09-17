package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AlphaWong/html2pdf/boot"
	"github.com/AlphaWong/html2pdf/utils"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		required.NoError(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusOK, rr.Code)
}

func TestPdfHandlerSuccess(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusOK, rr.Code)
}

func TestPdfHandlerFailBySameNameOptionValues(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Add fields
	// It must call before `CreateFormFile` as after it called, u cannot add fields
	// Reference: https://github.com/golang/go/blob/master/src/mime/multipart/writer.go#L145
	bodyWriter.WriteField("--footer-left", "[page] lalamove")
	bodyWriter.WriteField("--footer-left", "[page] lalamove")

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)

	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPdfHandlerSuccessByOptionValues(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Add fields
	// It must call before `CreateFormFile` as after it called, u cannot add fields
	// Reference: https://github.com/golang/go/blob/master/src/mime/multipart/writer.go#L145
	bodyWriter.WriteField("--footer-left", "[page] lalamove")

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)

	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusOK, rr.Code)
}

func TestPdfHandlerFailByTooBigFile(t *testing.T) {
	os.Setenv(utils.MaxSize, "0")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPdfHandlerFailByInvalidFileType(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/example.txt"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPdfHandlerFailByInvalidFormName(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("xxfile", filename)
	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPost, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusBadRequest, rr.Code)
}

func TestPdfHandlerFailByPutMethod(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "../simple/http2.html"
	utils.UploadPath = "../tmp/"
	utils.PdfPath = "../pdf/"

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	required.NoError(err)

	// open file handle
	fh, err := os.Open(filename)
	required.NoError(err)
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	required.NoError(err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	req, err := http.NewRequest(http.MethodPut, "/convert", bodyBuf)
	if err != nil {
		required.NoError(err)
	}
	req.Header.Set("Content-Type", contentType)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusMethodNotAllowed, rr.Code)
}
