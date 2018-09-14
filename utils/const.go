package utils

import (
	"github.com/segmentio/ksuid"
)

const (
	Port      = ":80"
	MaxSize   = "MAX_SIZE"
	SessionID = "SessionID"
)

var (
	MaxUploadSize int64
	UploadPath    = "./tmp/"
	PdfPath       = "./pdf/"
)

func GetUUID() string {
	return ksuid.New().String()
}
