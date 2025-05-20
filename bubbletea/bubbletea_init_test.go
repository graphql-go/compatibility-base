package bubbletea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleTeaInit(t *testing.T) {
	b := New(&Params{
		Models:    Models{NewChoicesModel(&ChoicesModelParams{})},
		BaseStyle: NewBaseStyle(),
	})

	assert.Nil(t, b.Init(), "unexpected non-nil result")
}
