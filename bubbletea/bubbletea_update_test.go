package bubbletea

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"

	customAssert "github.com/graphql-go/compatibility-base/assert"
)

func TestBubbleTeaUpdate(t *testing.T) {
	tests := []struct {
		subTestName      string
		initialBubbletea tea.Model
		updateParams     tea.Msg
		expectedModel    tea.Model
		expectedCmd      tea.Cmd
	}{
		{
			subTestName:      "Handles invalid tea key message",
			initialBubbletea: New(&Params{}),
			updateParams:     nil,
			expectedModel:    BubbleTea{},
			expectedCmd:      (tea.Cmd)(nil),
		},
		{
			subTestName:      "Handles ctrl+c tea key message",
			initialBubbletea: New(&Params{}),
			updateParams: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
			expectedModel: BubbleTea{},
			expectedCmd:   tea.Quit,
		},
		{
			subTestName: "Handles enter tea key message",
			initialBubbletea: New(&Params{
				Choices: []string{"test-choice-0"},
			}),
			updateParams: tea.KeyMsg{
				Type: tea.KeyEnter,
			},
			expectedModel: BubbleTea{
				cursor:  0,
				choices: []string{"test-choice-0"},
				choice:  "test-choice-0",
			},
			expectedCmd: tea.Quit,
		},
		{
			subTestName: "Handles down tea key message",
			initialBubbletea: New(&Params{
				Choices: []string{"test-choice-0", "test-choice-1"},
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("j"),
			},
			expectedModel: BubbleTea{
				cursor:  1,
				choices: []string{"test-choice-0", "test-choice-1"},
			},
			expectedCmd: nil,
		},
		{
			subTestName: "Handles cursor reset for down tea key message",
			initialBubbletea: New(&Params{
				Cursor:  1,
				Choices: []string{"test-choice-0", "test-choice-1"},
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("j"),
			},
			expectedModel: BubbleTea{
				cursor:  0,
				choices: []string{"test-choice-0", "test-choice-1"},
			},
			expectedCmd: nil,
		},
		{
			subTestName: "Handles up tea key message",
			initialBubbletea: New(&Params{
				Cursor:  1,
				Choices: []string{"test-choice-0", "test-choice-1"},
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("k"),
			},
			expectedModel: BubbleTea{
				cursor:  0,
				choices: []string{"test-choice-0", "test-choice-1"},
			},
			expectedCmd: nil,
		},
		{
			subTestName: "Handles cursor reset for up tea key message",
			initialBubbletea: New(&Params{
				Cursor:  0,
				Choices: []string{"test-choice-0"},
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("k"),
			},
			expectedModel: BubbleTea{
				cursor:  0,
				choices: []string{"test-choice-0"},
			},
			expectedCmd: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			model, cmd := tt.initialBubbletea.Update(tt.updateParams)

			assert.Equal(t, tt.expectedModel, model)
			customAssert.AssertFunc(t, tt.expectedCmd, cmd)
		})
	}
}
