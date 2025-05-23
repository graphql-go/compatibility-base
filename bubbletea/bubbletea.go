package bubbletea

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the component that wraps the `bubbletea` Model interface.
type Model interface {
	// Init initializes the model and returns the next tea command.
	Init() tea.Cmd

	// Update updates the model and returns the same model and the next tea command.
	Update(msg tea.Msg) (Model, tea.Cmd)

	// View returns the model view representation.
	View() string

	// Result returns the model result.
	Result() any

	// WithBaseStyle updates the model to use the given base style.
	WithBaseStyle(baseStyle lipgloss.Style)

	// Order returns the order of the model.
	Order() uint
}

// Models are the slice of Model interfaces.
type Models []Model

// `First` returns the first model of the `Models` slice.
func (m Models) First() Model {
	if len(m) == 1 {
		return m[0]
	}

	for _, model := range m {
		if model.Order() == 1 {
			return model
		}
	}

	return nil
}

// BubbleTea represents the CLI component that wraps the `bubbletea` library.
type BubbleTea struct {
	// baseStyle is the base styling of the BubbleTea component.
	baseStyle lipgloss.Style

	// currentModel is the current model of the BubbleTea component.
	currentModel Model

	// models are the models of the `BubbleTea` component.
	models Models
}

// `BubbleTeaResult` represents the `BubbleTea` component run result.
type BubbleTeaResult struct {
	// `ChoicesModelResult` is the result of the `ChoicesModel` component.
	ChoicesModelResult *ChoicesModelResult

	// `TableModelResult` is the result of the `TableModel` component.
	TableModelResult *TableModelResult
}

// Init is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Init() tea.Cmd {
	return b.currentModel.Init()
}

// Update is the `BubbleTea` method required for implementing the `Model` interface.
// Returns the `BubbleTea` struct and the next tea command(In case of `nil`, it indicates that the bubbletea program
// continues to work).
func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return b, nil
	}

	model, cmd := b.currentModel.Update(msg)

	switch keyMsg.String() {
	case "enter":
		nextModel := b.NextModel()
		if nextModel == nil {
			return b, cmd
		}

		b.currentModel = nextModel
		b.currentModel.WithBaseStyle(b.baseStyle)
		return b, nil
	default:
		b.currentModel = model
	}

	return b, cmd
}

// View is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) View() string {
	return b.currentModel.View()
}

// Run runs the `BubbleTea` component and returns its result.
func (b BubbleTea) Run() (*BubbleTeaResult, error) {
	teaProgram := tea.NewProgram(b)

	if _, err := teaProgram.Run(); err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	return b.Result()
}

// `Result` returns the `BubbleTea` component result.
func (b BubbleTea) Result() (*BubbleTeaResult, error) {
	r := &BubbleTeaResult{}

	for _, m := range b.models {
		switch m.(type) {
		case *ChoicesModel:
			result, ok := m.Result().(*ChoicesModelResult)
			if !ok {
				return nil, errors.New("unexpected type")
			}

			r.ChoicesModelResult = result
		case *TableModel:
			result, ok := m.Result().(*TableModelResult)
			if !ok {
				return nil, errors.New("unexpected type")
			}

			r.TableModelResult = result
		default:
			return nil, errors.New("unexpected type")
		}
	}

	return r, nil
}

// NextModel returns the next model, ordered by the order field.
func (b BubbleTea) NextModel() Model {
	currentOrder := b.currentModel.Order()
	nextOrder := currentOrder + 1

	for _, m := range b.models {
		if m.Order() == nextOrder {
			return m
		}
	}

	return nil
}

// Params represents the parameters for the `NewBubbleTea` function.
type Params struct {
	// Models are the models parameters for the `BubbleTea` component.
	Models Models

	// BaseStyle is the base styling parameter of the BubbleTea component.
	BaseStyle lipgloss.Style
}

// NewBaseStyle returns the default lipgloss base style.
func NewBaseStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("102")).
		Bold(false).
		PaddingTop(1).
		PaddingLeft(1).
		Width(124)
}

// New returns a new BubbleTea struct instance.
func New(p *Params) *BubbleTea {
	b := &BubbleTea{}

	b.baseStyle = p.BaseStyle
	b.models = p.Models
	b.currentModel = p.Models.First()
	b.currentModel.WithBaseStyle(p.BaseStyle)

	return b
}
