package utils

import (
	"github.com/segmentio/ksuid"
)

const (
	Port    = ":8000"
	MaxSize = "MAX_SIZE"
)

var (
	MaxUploadSize int64
	UploadPath    = "./tmp/"
	PdfPath       = "./pdf/"
)

func GetUUID() string {
	return ksuid.New().String()
}
