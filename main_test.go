package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AlphaWong/html2pdf/boot"
	"github.com/AlphaWong/html2pdf/handlers"
	"github.com/AlphaWong/html2pdf/utils"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	os.Setenv(utils.MaxSize, "20")
	boot.Init()
	var required = require.New(t)
	var filename = "./simple/people.html"

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
	handler := http.HandlerFunc(handlers.PdfHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	required.Equal(http.StatusOK, rr.Code)
}
