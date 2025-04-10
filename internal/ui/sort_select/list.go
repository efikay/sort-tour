package sort_select

import (
	"sort-tour/internal/entities/sort"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	lip "github.com/charmbracelet/lipgloss"
)

var docStyle = lip.NewStyle().Margin(1, 2)

type listModel struct {
	list list.Model
}

// Init implements tea.Model.
func (m listModel) Init() tea.Cmd {
	return nil
}

// View implements tea.Model.
func (m listModel) View() string {
	return docStyle.Render(m.list.View())
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()

		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func newListModel() listModel {
	m := listModel{
		list: list.New(sort.GetSortInfoListItems(), list.NewDefaultDelegate(), 0, 0),
	}
	m.list.Title = "Choose sorting algorithm"

	return m
}
