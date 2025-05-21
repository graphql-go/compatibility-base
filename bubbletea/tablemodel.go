package bubbletea

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// `TableModel` represents the component that implements the `Model` interface.
type TableModel struct {
	// baseStyle is the base styling for the `table` model.
	baseStyle lipgloss.Style

	// table is the `bubbletea` table model.
	table table.Model

	// order is the order of the model
	order uint
}

// `Init` is the `TableModel` method required for implementing the `Model` interface.
// Initializes the `TableModel` component and returns a `bubbletea` command.
func (tm TableModel) Init() tea.Cmd {
	return nil
}

// `Update` is the `TableModel` method required for implementing the `Model` interface.
// Updates the `TableModel` component, handles the given message updating the internal state.
// Returns the current `TableModel` and the resolved command.
// TODO(@chris-ramon): Implement the `tea.Msg` handlers.
func (tm *TableModel) Update(msg tea.Msg) (Model, tea.Cmd) { //nolint:golint,ireturn
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
	return tm.baseStyle.Render(tm.table.View()) + "\n"
}

// `TableModelResult` represents the `TableModel` run method result.
type TableModelResult struct {
}

// `Run` is the `TableModel` method required for implementing the `Model` interface.
// Runs the `TableModel` component and returns its result.
func (tm TableModel) Run(model any) (any, error) {
	result := &TableModelResult{}

	if m, ok := model.(TableModel); ok {
		return m, nil
	}

	return result, nil
}

// `WithBaseStyle` updates the `TableModel` component to use the given base style.
func (tm *TableModel) WithBaseStyle(baseStyle lipgloss.Style) {
	tm.baseStyle = baseStyle
}

func (tm *TableModel) Order() uint {
	return tm.order
}

// `TableModelParams` represents the parameters component for the `NewTableModel` function.
type TableModelParams struct {
	// BaseStyle is the base styling parameter.
	BaseStyle lipgloss.Style

	// Order is the order parameter.
	Order uint
}

// `NewTableModel` returns a pointer to a `TableModel`.
func NewTableModel(p *TableModelParams) *TableModel {
	widthColumn := 15

	columns := []table.Column{
		{Title: "Metric", Width: 43},
		{Title: "Spec: https://github.com/graphql/graphql-js", Width: widthColumn},
		{Title: "Impl: https://github.com/graphql-go/graphql", Width: widthColumn},
		{Title: "Diff Ratio", Width: widthColumn},
		{Title: "Max Diff", Width: widthColumn},
		{Title: "Result", Width: widthColumn},
	}

	rows := []table.Row{
		{"GitHub:", "", "", "", "", ""},
		{"License", "MIT", "MIT", "0%", "0%", "✅"},
		{"Number Of Stars", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Issues Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Issues Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Pull Requests Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Pull Requests Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Forks", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Last Commit Date", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Contributors", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"GraphQL Compatibility Keywords:", "", "", "", "", ""},
		{"Number Of Comments Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"Number Of Comments Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		{"GraphQL:", "", "", "", "", ""},
		{"Specification Version", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(17),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#a8a7a3")).
		Bold(false)

	t.SetStyles(s)

	return &TableModel{
		baseStyle: p.BaseStyle,
		table:     t,
		order:     p.Order,
	}
}
