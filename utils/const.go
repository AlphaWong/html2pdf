package utils

import (
	"github.com/segmentio/ksuid"
)

const (
	Port       = ":8000"
	UploadPath = "./tmp/"
	PdfPath    = "./pdf/"

	MaxSize = "MAX_SIZE"
)

var (
	MaxUploadSize int64
)

func GetUUID() string {
	return ksuid.New().String()
}
