package bubbletea

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ChoicesModel represents the CLI component that wraps the `bubbletea` library.
type ChoicesModel struct {
	// cursor is the reference of the current CLI choice.
	cursor int

	// choice is the current CLI choice.
	choice string

	// choices is the slice of CLI choices.
	choices []string

	// ui is the UI of the CLI.
	ui ChoicesModelUI

	// baseStyle is the base styling of the BubbleTea component.
	baseStyle lipgloss.Style
}

// ChoicesModelUI represents the UI struct for the `ChoicesModel` component.
type ChoicesModelUI struct {
	// header is the UI header text.
	header string
}

// Init is the `BubbleTea` method required for implementing the `Model` interface.
func (b ChoicesModel) Init() tea.Cmd {
	return nil
}

// Update is the `BubbleTea` method required for implementing the `Model` interface.
func (b ChoicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:golint,ireturn
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return b, nil
	}

	switch keyMsg.String() {
	case "ctrl+c", "q", "esc":
		return b, tea.Quit

	case "enter":
		if len(b.choices) > 0 {
			b.choice = b.choices[b.cursor]
		}

		return b, tea.Quit

	case "down", "j":
		b.cursor++
		if b.cursor >= len(b.choices) {
			b.cursor = 0
		}

	case "up", "k":
		b.cursor--
		if b.cursor < 0 {
			b.cursor = len(b.choices) - 1
		}
	}

	return b, nil
}

// View is the `BubbleTea` method required for implementing the `Model` interface.
func (b ChoicesModel) View() string {
	s := strings.Builder{}
	s.WriteString(b.ui.header)
	s.WriteString("")

	for i := range b.choices {
		if b.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		choice := b.choices[i]
		s.WriteString(choice)

		s.WriteString("\n")
	}

	endingMessage := "\n(press q to quit)\n"
	s.WriteString(endingMessage)

	choicesModelView := s.String()

	return b.baseStyle.Render(choicesModelView) + "\n"
}

// ChoicesModelResult represents the result of the run method.
type ChoicesModelResult struct {
	// Choice is the option chosen.
	Choice string
}

// Run runs the `ChoicesModel` component and returns its result.
func (b ChoicesModel) Run(model any) (any, error) {
	result := &ChoicesModelResult{}

	if model, ok := model.(ChoicesModel); ok && model.choice != "" {
		result.Choice = model.choice
	}

	return result, nil
}

func (b *ChoicesModel) WithBaseStyle(baseStyle lipgloss.Style) {
	b.baseStyle = baseStyle
}

// ChoicesModelParams represents the parameters struct for the `NewChoicesModel` function.
type ChoicesModelParams struct {
	// Choice is the current CLI choice.
	Choice string

	// Choices is the slice of options available.
	Choices []string

	// Cursor is the reference of the current CLI choice.
	Cursor int

	// UI is the user interface parameters.
	UI ChoicesModelUIParams

	// BaseStyle is the base styling parameter of the BubbleTea component.
	BaseStyle lipgloss.Style
}

// ChoicesModelUIParams represents the UI parameters for the `NewChoicesModel` function parameters.
type ChoicesModelUIParams struct {
	// Header is the UI header text.
	Header string
}

// NewChoicesModel returns a pointer for the `ChoicesModel`.
func NewChoicesModel(p *ChoicesModelParams) *ChoicesModel {
	return &ChoicesModel{
		choice:  p.Choice,
		choices: p.Choices,
		cursor:  p.Cursor,
		ui: ChoicesModelUI{
			header: p.UI.Header,
		},
		baseStyle: p.BaseStyle,
	}
}
