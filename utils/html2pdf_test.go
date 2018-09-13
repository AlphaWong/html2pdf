package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFormValues(t *testing.T) {
	var required = require.New(t)
	formValues := map[string][]string{
		"a": []string{"1", "2", "3"},
		"b": []string{"4", "5", "6"},
	}
	var r = ParseFormValues(formValues)
	required.Contains(r, "a 1,2,3")
	required.Contains(r, "b 4,5,6")
}
