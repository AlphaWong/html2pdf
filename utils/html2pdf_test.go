package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFormValuesSuccess(t *testing.T) {
	var required = require.New(t)
	formValues := map[string][]string{
		"a": []string{"1,2,3"},
		"b": []string{"4,5,6"},
	}
	var r, err = ParseFormValues(formValues)
	required.NoError(err)
	required.Contains(r, "a")
	required.Contains(r, "1,2,3")
	required.Contains(r, "b")
	required.Contains(r, "4,5,6")
}

func TestParseFormValuesFailByMultipleValues(t *testing.T) {
	var required = require.New(t)
	formValues := map[string][]string{
		"a": []string{"1,2,3", "9,9,9"},
		"b": []string{"4,5,6"},
	}
	var r, err = ParseFormValues(formValues)
	required.Error(err)
	required.Equal([]string{}, r)
}
