package bubbletea

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// BubbleTea represents the CLI component that wraps the `bubbletea` library.
type BubbleTea struct {
	// cursor is the reference of the current CLI choice.
	cursor int

	// choice is the current CLI choice.
	choice string

	// choices is the slice of CLI choices.
	choices []string

	// ui is the UI of the CLI.
	ui UI

	// table is the bubbletea table model.
	table table.Model
}

// UI represents the UI struct for the `BubbleTea` component.
type UI struct {
	// header is the UI header text.
	header string
}

// Init is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Init() tea.Cmd {
	return nil
}

// Update is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:golint,ireturn
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
func (b BubbleTea) View() string {
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

	return s.String()
}

// RunResult represents the result of the run method.
type RunResult struct {
	// Choice is the option chosen.
	Choice string
}

// Params represents the parameters struct for the new method.
type Params struct {
	// Choices is the slice of options available.
	Choices []string

	// Cursor is the reference of the current CLI choice.
	Cursor int

	// UI is the user interface parameters.
	UI UIParams
}

// UIParams represents the UI parameters for the new method parameters.
type UIParams struct {
	// Header is the UI header text.
	Header string
}

// New returns a pointer for the `BubbleTea` component.
func New(p *Params) *BubbleTea {
	return &BubbleTea{
		choices: p.Choices,
		cursor:  p.Cursor,
		ui: UI{
			header: p.UI.Header,
		},
	}
}

// Run runs the `BubbleTea` component and returns its result.
func (b BubbleTea) Run() (*RunResult, error) {
	teaProgram := tea.NewProgram(b)

	m, err := teaProgram.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	result := &RunResult{}

	if m, ok := m.(BubbleTea); ok && m.choice != "" {
		result.Choice = m.choice
	}

	return result, nil
}
