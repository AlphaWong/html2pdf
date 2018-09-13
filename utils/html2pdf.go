package utils

import (
	"fmt"
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

func ParseFormValues(formValues map[string][]string) ([]string, error) {
	// The orderfing of the formValues is not always the same
	// User disallow input same name of the key
	//
	// For example
	// curl -X POST \
	// https://httpbin.org/post \
	// -H 'Cache-Control: no-cache' \
	// -H 'Postman-Token: d820d5d1-aa58-4d9c-9f43-f5636e67dcbe' \
	// -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
	// -F 'a=1,2,3' \
	// -F 'b=4,5,6' \
	// -F a=4
	// key `a` has been insert two times.
	//
	// However, Go do not allow it in `CreateFormField`
	// Reference: https://github.com/golang/go/blob/master/src/mime/multipart/writer.go#L145
	var ss = []string{}
	for k, v := range formValues {
		if len(v) > 1 {
			return []string{}, fmt.Errorf("key %s has been inputed more than one times", k)
		}
		ss = append(ss, k, v[0])
	}
	return ss, nil
}
