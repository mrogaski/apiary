package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/go-template"
)

func Test_Greeting(t *testing.T) {
	t.Parallel()

	got, err := template.Greeting()

	if assert.NoError(t, err) {
		assert.Equal(t, got, "Hello, world!")
	}
}
