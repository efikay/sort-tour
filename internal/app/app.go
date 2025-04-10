package app

import (
	"fmt"
	"log"
	"sort-tour/internal/ui/sort_select"

	tea "github.com/charmbracelet/bubbletea"
)

type appWindow uint

const (
	chooseSortWindow appWindow = iota
	sortWindow
)

type appModel struct {
	activeWindow appWindow

	sortSelectWindow sort_select.Window
}

// Init implements tea.Model.
func (m appModel) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case sort_select.WindowModelQuitMessage:
		return m, tea.Quit
	case sort_select.WindowModelChooseSortMessage:
		sortName := msg.SortName

		panic(fmt.Sprintf("Chosen sort name: %s", sortName))
		// TODO: Create sort window for SortName sort
	}

	switch m.activeWindow {
	case chooseSortWindow:
		m.sortSelectWindow, cmd = m.sortSelectWindow.Update(msg)
		return m, cmd
	case sortWindow:
		// TODO: Make sort window
	}

	panic(fmt.Sprintf("unexpected app.appWindow: %#v", m.activeWindow))
}

// View implements tea.Model.
func (m appModel) View() string {
	switch m.activeWindow {
	case chooseSortWindow:
		return m.sortSelectWindow.View()
	case sortWindow:
		// TODO: Make sort window
	}

	panic(fmt.Sprintf("unexpected app.appWindow: %#v", m.activeWindow))
}

func Run() {
	p := tea.NewProgram(appModel{
		activeWindow:     chooseSortWindow,
		sortSelectWindow: sort_select.NewWindowModel(),
	}, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
