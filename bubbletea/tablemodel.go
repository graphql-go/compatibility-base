package bubbletea

import (
	"fmt"

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
	endingMessage := "\n(press enter to continue)\n"
	view := tm.baseStyle.Render(tm.table.View()) + "\n"
	return fmt.Sprintf("%s%s", view, endingMessage)
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

	// Headers are the slice of table headers.
	Headers []TableHeader

	// Rows are the slice of table rows.
	Rows [][]string
}

// TableHeader represents the table header.
type TableHeader struct {
	// Title is the title of the table header.
	Title string

	// Width is the styling width of the table header.
	Width uint
}

// `NewTableModel` returns a pointer to a `TableModel`.
func NewTableModel(p *TableModelParams) *TableModel {
	columnsHeaders := []table.Column{}

	for _, h := range p.Headers {
		header := table.Column{Title: h.Title, Width: int(h.Width)}
		columnsHeaders = append(columnsHeaders, header)
	}

	dataRows := []table.Row{}

	for _, r := range p.Rows {
		dataRows = append(dataRows, table.Row(r))
	}

	t := table.New(
		table.WithColumns(columnsHeaders),
		table.WithRows(dataRows),
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
