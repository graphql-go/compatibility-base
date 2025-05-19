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
		expectedModel    *BubbleTea
		expectedCmd      tea.Cmd
	}{
		{
			subTestName: "Handles invalid tea key message",
			initialBubbletea: New(&Params{
				Model:     NewChoicesModel(&ChoicesModelParams{}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: nil,
			expectedModel: New(&Params{
				Model:     NewChoicesModel(&ChoicesModelParams{}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: (tea.Cmd)(nil),
		},
		{
			subTestName: "Handles ctrl+c tea key message",
			initialBubbletea: New(&Params{
				Model:     NewChoicesModel(&ChoicesModelParams{}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type: tea.KeyCtrlC,
			},
			expectedModel: New(&Params{
				Model:     NewChoicesModel(&ChoicesModelParams{}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: tea.Quit,
		},
		{
			subTestName: "Handles enter tea key message",
			initialBubbletea: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Choices: []string{"test-choice-0"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type: tea.KeyEnter,
			},
			expectedModel: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:    0,
					Choices:   []string{"test-choice-0"},
					Choice:    "test-choice-0",
					BaseStyle: NewBaseStyle(),
				}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: tea.Quit,
		},
		{
			subTestName: "Handles down tea key message",
			initialBubbletea: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("j"),
			},
			expectedModel: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  1,
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: nil,
		},
		{
			subTestName: "Handles cursor reset for down tea key message",
			initialBubbletea: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  1,
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("j"),
			},
			expectedModel: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  0,
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: nil,
		},
		{
			subTestName: "Handles up tea key message",
			initialBubbletea: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  1,
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("k"),
			},
			expectedModel: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  0,
					Choices: []string{"test-choice-0", "test-choice-1"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: nil,
		},
		{
			subTestName: "Handles cursor reset for up tea key message",
			initialBubbletea: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  0,
					Choices: []string{"test-choice-0"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			updateParams: tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune("k"),
			},
			expectedModel: New(&Params{
				Model: NewChoicesModel(&ChoicesModelParams{
					Cursor:  0,
					Choices: []string{"test-choice-0"},
				}),
				BaseStyle: NewBaseStyle(),
			}),
			expectedCmd: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.subTestName, func(t *testing.T) {
			model, cmd := tt.initialBubbletea.Update(tt.updateParams)

			assert.Equal(t, *tt.expectedModel, model)
			customAssert.AssertFunc(t, tt.expectedCmd, cmd)
		})
	}
}
