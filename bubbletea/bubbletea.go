package bubbletea

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the component that wraps the `bubbletea` Model interface.
type Model interface {
	// tea.Model is the base interface.
	tea.Model

	// Run runs and returns its result.
	Run(modelResult any) (any, error)

	// WithBaseStyle updates the model to use the given base style.
	WithBaseStyle(baseStyle lipgloss.Style)
}

// BubbleTea represents the CLI component that wraps the `bubbletea` library.
type BubbleTea struct {
	// baseStyle is the base styling of the BubbleTea component.
	baseStyle lipgloss.Style

	// currentModel is the current model of the BubbleTea component.
	currentModel Model
}

// Init is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Init() tea.Cmd {
	return b.currentModel.Init()
}

// Update is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_, cmd := b.currentModel.Update(msg)
	return b, cmd
}

// View is the `BubbleTea` method required for implementing the `Model` interface.
func (b BubbleTea) View() string {
	return b.currentModel.View()
}

// Run runs the `BubbleTea` component and returns its result.
func (b BubbleTea) Run() (any, error) {
	teaProgram := tea.NewProgram(b)

	m, err := teaProgram.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	return b.currentModel.Run(m)
}

// Params represents the parameters for the `NewBubbleTea` function.
type Params struct {
	// Model is the model parameter of the BubbleTea component.
	Model Model

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
		Width(80)
}

// New returns a new BubbleTea struct instance.
func New(p *Params) *BubbleTea {
	b := &BubbleTea{}

	b.baseStyle = p.BaseStyle
	b.currentModel = p.Model
	b.currentModel.WithBaseStyle(p.BaseStyle)

	return b
}
