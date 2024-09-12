package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Deck struct {
	name      string
	review    int
	total     int
	createdAt string
}

type model struct {
	decks       []Deck
	hoveredDeck int
}

var initialModel = model{
	decks: []Deck{
		{
			"Spanish 🇪",
			7,
			23,
			"January",
		},
		{
			"German 🇩",
			13,
			109,
			"February",
		},
	},
	hoveredDeck: 0,
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.hoveredDeck > 0 {
				m.hoveredDeck--
			}

		case "down", "j":
			if m.hoveredDeck < len(m.decks)-1 {
				m.hoveredDeck++
			}

		case "enter", " ":
		  fmt.Println("Selected")
    }
	}

	return m, nil
}

func (m model) View() string {
  var style = lipgloss.NewStyle().
    BorderStyle(lipgloss.RoundedBorder()).
    Foreground(lipgloss.Color("#FAFAFA")).
    Background(lipgloss.Color("#7D56F4")).
    PaddingLeft(4).
    Width(22)

    return style.Render("Hello Kitty")
}

func main() {
	fmt.Println("Hello this is my new golang project!")
  p := tea.NewProgram(initialModel)
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error")
    os.Exit(1)
  }
}
