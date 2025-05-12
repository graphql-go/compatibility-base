package bubbletea

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BubbleTea represents the CLI component that wraps the `bubbletea` library.
type BubbleTea struct {
	// baseStyle is the base styling of the BubbleTea component.
	baseStyle lipgloss.Style

	// currentModel is the current model of the BubbleTea component.
	currentModel tea.Model
}

// Init is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Init() tea.Cmd {
	return nil
}

// Update is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return b, nil
	}

	switch keyMsg.String() {
	case "ctrl+c", "q", "esc":
		return b, tea.Quit
	}

	return b, nil
}

// View is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) View() string {
	return b.baseStyle.Render(b.currentModel.View()) + "\n"
}

// Run runs the `BubbleTea` component and returns its result.
func (b BubbleTea) Run() (*RunResult, error) {
	teaProgram := tea.NewProgram(b)

	_, err := teaProgram.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	result := &RunResult{}

	// if m, ok := m.(BubbleTea); ok && m.choice != "" {
	// 	result.Choice = m.choice
	// }

	return result, nil
}

// RunResult represents the result of the run method.
type RunResult struct {
	// Choice is the option chosen.
	Choice string
}

// Params represents the parameters for the `NewBubbleTea` function.
type Params struct {
	// Model is the model parameter of the BubbleTea component.
	Model tea.Model
}

// NewBubbleTea returns a new BubbleTea struct instance.
func New(p *Params) *BubbleTea {
	var baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))

	b := &BubbleTea{}

	b.baseStyle = baseStyle
	b.currentModel = p.Model

	return b
}

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

	// table is the bubbletea table model.
	// TODO(@chris-ramon): Make it available through multiple model support.
	table table.Model
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

	return s.String()
}

// ChoicesModelParams represents the parameters struct for the `NewChoicesModel` function.
type ChoicesModelParams struct {
	// Choices is the slice of options available.
	Choices []string

	// Cursor is the reference of the current CLI choice.
	Cursor int

	// UI is the user interface parameters.
	UI ChoicesModelUIParams
}

// ChoicesModelUIParams represents the UI parameters for the `NewChoicesModel` function parameters.
type ChoicesModelUIParams struct {
	// Header is the UI header text.
	Header string
}

// NewChoicesModel returns a pointer for the `ChoicesModel`.
func NewChoicesModel(p *ChoicesModelParams) *ChoicesModel {
	return &ChoicesModel{
		choices: p.Choices,
		cursor:  p.Cursor,
		ui: ChoicesModelUI{
			header: p.UI.Header,
		},
	}
}
