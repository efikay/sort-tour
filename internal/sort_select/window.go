package sort_select

import (
	"sort-tour/internal/entities/sort"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Window struct {
	sortList listModel
}

func (m Window) View() string {
	return m.sortList.View()
}

func (m Window) Update(msg tea.Msg) (Window, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, func() tea.Msg {
				return WindowModelQuitMessage{}
			}
		case "enter":
			is_filtering := m.sortList.list.FilterState() == list.Filtering
			if !is_filtering {
				selected_item, ok := m.sortList.list.SelectedItem().(sort.ShortSortInfo)

				if ok {
					return m, func() tea.Msg {
						return WindowModelChooseSortMessage{
							SortName: selected_item.Title(),
						}
					}
				}
			}

		}
	}

	m.sortList, cmd = m.sortList.Update(msg)

	return m, cmd
}

func NewWindowModel() Window {
	return Window{
		sortList: newListModel(),
	}
}

type WindowModelQuitMessage struct{}
type WindowModelChooseSortMessage struct {
	SortName string
}
