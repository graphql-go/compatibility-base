package bubbletea

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
)

func TestBubbleTeaView(t *testing.T) {
	tests := []struct {
		subTestName      string
		initialBubbletea tea.Model
		expectedView     string
	}{
		{
			subTestName: "Handles success view result",
			initialBubbletea: New(&Params{
				Models: Models{NewChoicesModel(&ChoicesModelParams{
					Choices: []string{"test-choice-0"},
					UI: ChoicesModelUIParams{
						Header: "test-header: \n",
					},
				})},
				BaseStyle: NewBaseStyle(),
			}),
			expectedView: "┌────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n│                                                                                                                            │\n│ test-header:                                                                                                               │\n│ (•) test-choice-0                                                                                                          │\n│                                                                                                                            │\n│ (press q to quit)                                                                                                          │\n│                                                                                                                            │\n└────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘\n",
		},
		{
			subTestName: "Handles success view result with multiple choices",
			initialBubbletea: New(&Params{
				Models: Models{NewChoicesModel(&ChoicesModelParams{
					Choices: []string{"test-choice-0", "test-choice-1"},
					UI: ChoicesModelUIParams{
						Header: "test-header: \n",
					},
				})},
				BaseStyle: NewBaseStyle(),
			}),
			expectedView: "┌────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n│                                                                                                                            │\n│ test-header:                                                                                                               │\n│ (•) test-choice-0                                                                                                          │\n│ ( ) test-choice-1                                                                                                          │\n│                                                                                                                            │\n│ (press q to quit)                                                                                                          │\n│                                                                                                                            │\n└────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			view := tt.initialBubbletea.View()

			assert.Equal(t, tt.expectedView, view)
		})
	}
}
