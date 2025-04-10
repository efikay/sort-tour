package choose_sort

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type WindowModel struct {
	sortList listModel
}

func (m WindowModel) View() string {
	return m.sortList.View()
}

func (m WindowModel) Update(msg tea.Msg) (WindowModel, tea.Cmd) {
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
				selected_item, ok := m.sortList.list.SelectedItem().(item)

				if ok {
					return m, func() tea.Msg {
						return WindowModelChooseSortMessage{
							SortName: selected_item.title,
						}
					}
				}
			}

		}
	}

	m.sortList, cmd = m.sortList.Update(msg)

	return m, cmd
}

func NewWindowModel() WindowModel {
	return WindowModel{
		sortList: newListModel(),
	}
}

type WindowModelQuitMessage struct{}
type WindowModelChooseSortMessage struct {
	SortName string
}
