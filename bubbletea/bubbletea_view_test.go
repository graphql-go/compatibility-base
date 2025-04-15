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
				Choices: []string{"test-choice-0"},
				UI: UIParams{
					Header: "test-header: \n",
				},
			}),
			expectedView: `test-header: 
(•) test-choice-0

(press q to quit)
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			view := tt.initialBubbletea.View()

			assert.Equal(t, tt.expectedView, view)
		})
	}
}
