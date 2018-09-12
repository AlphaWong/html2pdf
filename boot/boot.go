package boot

import (
	"log"
	"os"
	"strconv"

	"github.com/AlphaWong/html2pdf/utils"
)

func Init() {
	var maxSize, err = strconv.ParseInt(os.Getenv(utils.MaxSize), 10, 64)
	if err != nil {
		log.Fatal("Invalid MAX_SIZE env")
	}
	utils.MaxUploadSize = maxSize * 1024 * 1024 // MB as unit
}
