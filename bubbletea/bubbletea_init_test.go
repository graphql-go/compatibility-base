package bubbletea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleTeaInit(t *testing.T) {
	b := BubbleTea{}

	assert.Nil(t, b.Init(), "unexpected non-nil result")
}
