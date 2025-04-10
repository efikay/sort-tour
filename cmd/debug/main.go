package debug

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	Main()
}

func Main() {
	items := []list.Item{
		item{title: "Bubble Sort", desc: "Repeatedly swaps adjacent elements if they're in the wrong order"},
		item{title: "Selection Sort", desc: "Selects the smallest element and swaps it into place each iteration"},
		item{title: "Insertion Sort", desc: "Builds the sorted array one element at a time by inserting each item into its correct position"},
		item{title: "Merge Sort", desc: "Divides the array into halves, recursively sorts, then merges them back together"},
		item{title: "Quick Sort", desc: "Picks a 'pivot' element and partitions the array around it, recursively sorting sub-arrays"},
		item{title: "Heap Sort", desc: "Converts the array into a max-heap to repeatedly extract the largest element"},
		item{title: "Counting Sort", desc: "Counts occurrences of each value (works only for integer ranges)"},
		item{title: "Radix Sort", desc: "Sorts numbers digit-by-digit (from least to most significant)"},
		item{title: "Bucket Sort", desc: "Distributes elements into buckets, sorts each bucket, then concatenates"},
		item{title: "Shell Sort", desc: "An optimized insertion sort that compares elements spaced apart"},
		item{title: "Tim Sort", desc: "Hybrid of merge sort and insertion sort (used in Python and Java)"},
		item{title: "Comb Sort", desc: "Improves bubble sort by comparing elements spaced apart (shrinking gap)"},
		item{title: "Cycle Sort", desc: "Minimizes writes by cycling elements to their correct positions"},
		item{title: "Cocktail Shaker Sort", desc: "Bidirectional bubble sort that alternates directions"},
		item{title: "Gnome Sort", desc: "Similar to insertion sort but swaps backward like a garden gnome"},
		item{title: "Bogo Sort", desc: "Randomly shuffles the array until it's sorted (theoretically terrible)"},
		item{title: "Stooge Sort", desc: "Recursively sorts first 2/3, last 2/3, then first 2/3 again"},
		item{title: "Pancake Sort", desc: "Flips subarrays to move the largest element to its position"},
		item{title: "Bitonic Sort", desc: "Works best on power-of-two-sized arrays (parallelizable)"},
		item{title: "Tree Sort", desc: "Uses a binary search tree to sort elements"},
		item{title: "Cube Sort", desc: "Parallel version of bubble sort for 3D data"},
		item{title: "Odd-Even Sort", desc: "Compares odd/even indexed pairs (parallelizable)"},
		item{title: "Strand Sort", desc: "Extracts sorted 'strands' from the list and merges them"},
		item{title: "Intro Sort", desc: "Hybrid of quick sort, heap sort, and insertion sort"},
		item{title: "Smooth Sort", desc: "Adaptive variant of heap sort with near-O(n) best-case performanc"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Sorting algorithms"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
