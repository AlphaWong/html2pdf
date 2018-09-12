package utils

import (
	"github.com/segmentio/ksuid"
)

const (
	Port       = ":8000"
	UploadPath = "./tmp/"
	PdfPath    = "./pdf/"
)

var (
	MaxUploadSize int64
)

func GetUUID() string {
	return ksuid.New().String()
}
