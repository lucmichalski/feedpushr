package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ncarlier/feedpushr/v3/pkg/strcase"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		value    string
		expected string
	}{
		{"hello-world", "hello_world"},
		{"helloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"Hello/_World", "hello__world"},
		{"Hello/world", "hello_world"},
	}
	for _, tc := range testCases {
		value := strcase.ToSnake(tc.value)
		assert.Equal(t, tc.expected, value)
	}
}
