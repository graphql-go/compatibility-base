package bubbletea

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// `TableModel` represents the component that implements the `Model` interface.
type TableModel struct {
	// table is the `bubbletea` table model.
	table table.Model

	// baseStyle is the base styling for the `table` model.
	baseStyle lipgloss.Style
}

// `Init` is the `TableModel` method required for implementing the `Model` interface.
// Initializes the `TableModel` component and returns a `bubbletea` command.
func (tb TableModel) Init() tea.Cmd {
	return nil
}

// `Update` is the `TableModel` method required for implementing the `Model` interface.
// Updates the `TableModel` component, handles the given message updating the internal state.
// Returns the current `TableModel` and the resolved command.
func (tm TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:golint,ireturn
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return tm, nil
	}

	switch keyMsg.String() {
	case "ctrl+c", "q", "esc":
		return tm, tea.Quit

	case "enter":
		return tm, tea.Quit

	case "down", "j":
		return tm, tea.Quit

	case "up", "k":
		return tm, tea.Quit
	}

	return tm, nil
}

// `View` is the `TableModel` method required for implementing the `Model` interface.
// View renders the `TableModel` using the base style and returns the results.
func (tm TableModel) View() string {
	return tm.baseStyle.Render("") + "\n"
}

// `Run` is the `TableModel` method required for implementing the `Model` interface.
// Runs the `TableModel` component and returns its result.
func (tm TableModel) Run(model any) (any, error) {
	result := &ChoicesModelResult{}

	if m, ok := model.(TableModel); ok {
		return m, nil
	}

	return result, nil
}

// `WithBaseStyle` updates the `TableModel` component to use the given base style.
func (tm *TableModel) WithBaseStyle(baseStyle lipgloss.Style) {
	tm.baseStyle = baseStyle
}

// `TableModelParams` represents the parameters component for the `NewTableModel` function.
type TableModelParams struct {
	// BaseStyle is the base styling parameter.
	BaseStyle lipgloss.Style
}

// `NewTableModel` returns a pointer to a `TableModel`.
func NewTableModel(p *TableModelParams) *TableModel {
	return &TableModel{
		baseStyle: p.BaseStyle,
	}
}
