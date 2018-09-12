package utils

import (
	"os/exec"
	"strings"
)

type Converter interface {
	ConvertHtml2Pdf() error
}

type Html2Pdf struct {
	InFilePath  string
	OutFilePath string
}

type ConverterParam struct {
	InFilePath  string
	OutFilePath string
}

var _ Converter = (*Html2Pdf)(nil)

const wkhtmltopdf = "wkhtmltopdf"

func NewConverter(c *ConverterParam) Converter {
	return &Html2Pdf{
		InFilePath:  c.InFilePath,
		OutFilePath: c.OutFilePath,
	}
}

func (h *Html2Pdf) ConvertHtml2Pdf() error {
	var cmd = exec.Command(wkhtmltopdf, h.InFilePath, h.OutFilePath)
	var out strings.Builder
	cmd.Stderr = &out
	err := cmd.Run()
	s := out.String()
	ss := strings.Split(s, "\n")
	for i := len(ss) - 1; i > -1; i-- {
		if strings.Contains(ss[i], "Done") {
			return nil
		} else if strings.Contains(ss[i], "Printing pages") {
			return err
		}
	}
	return err
}
