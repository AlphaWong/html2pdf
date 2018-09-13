package utils

import (
	"log"
	"os/exec"
	"strings"
)

type Converter interface {
	ConvertHtml2Pdf() error
}

type Html2Pdf struct {
	InFilePath  string
	OutFilePath string
	Options     []string
}

type ConverterParam struct {
	InFilePath  string
	OutFilePath string
	Options     []string
}

var _ Converter = (*Html2Pdf)(nil)

const wkhtmltopdf = "wkhtmltopdf"

func NewConverter(c *ConverterParam) Converter {
	return &Html2Pdf{
		InFilePath:  c.InFilePath,
		OutFilePath: c.OutFilePath,
		Options:     c.Options,
	}
}

func (h *Html2Pdf) ConvertHtml2Pdf() error {
	// Create a cmd with options
	var cmd = exec.Command(wkhtmltopdf, h.Options...)
	// append the input path and output path
	cmd.Args = append(cmd.Args, h.InFilePath, h.OutFilePath)
	log.Println(cmd.Args)
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

func ParseFormValues(formValues map[string][]string) []string {
	// The orderfing of the formValues is not always the same
	var ss = []string{}
	for k, v := range formValues {
		ss = append(ss, k, strings.Join(v, ","))
	}
	return ss
}
