package bubbletea

import (
	"reflect"
	"runtime"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func TestBubbleTeaInit(t *testing.T) {
	b := BubbleTea{}

	if cmd := b.Init(); cmd != nil {
		t.Fatalf("expected: nil, got: %v", cmd)
	}
}

func TestBubbleTeaUpdate(t *testing.T) {
	b := BubbleTea{}

	tests := []struct {
		subTestName   string
		updateParams  tea.Msg
		expectedModel tea.Model
		expectedCmd   tea.Cmd
	}{
		{
			subTestName:   "Handles invalid tea key message",
			updateParams:  nil,
			expectedModel: b,
			expectedCmd:   (tea.Cmd)(nil),
		},
		{
			subTestName: "Handles ctrl+c tea key message",
			updateParams: tea.KeyMsg{
				Type:  tea.KeyCtrlC,
				Runes: []rune{'q'},
			},
			expectedModel: b,
			expectedCmd:   tea.Quit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			model, cmd := b.Update(tt.updateParams)

			assert.Equal(t, tt.expectedModel, model)
			assertFunc(t, tt.expectedCmd, cmd)
		})
	}
}

// assertFunc asserts the given two functions.
func assertFunc(t *testing.T, expected any, actual any) {
	expectedFunc := runtime.FuncForPC(reflect.ValueOf(expected).Pointer()).Name()
	actualFunc := runtime.FuncForPC(reflect.ValueOf(actual).Pointer()).Name()

	assert.Equal(t, expectedFunc, actualFunc)
}
