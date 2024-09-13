package main

import (
	"Szuhaydv/golang_proj/styles"

	"github.com/charmbracelet/lipgloss"
	"os"
  "fmt"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type Deck struct {
	name      string
	review    string
	total     string
	createdAt string
}

type model struct {
	decks       []Deck
	hoveredDeck int
}

func initialModel() model {
  return model{
    decks: []Deck{
      {
        "Spanish 🇪🇸",
        "15",
        "95",
        "2012-12-24",
      },
      {
        "German 🇩🇪",
        "34",
        "128",
        "2011-05-14",
      },
      {
        "English 🇬🇧",
        "9",
        "36",
        "2009-07-18",
      },
    },
    hoveredDeck: 0,
  }
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
  var isDecksEmpty bool
	if len(m.decks) == 0 {
    isDecksEmpty = true
  }
  header := styles.Header(isDecksEmpty)

  rows := []string{header}
  for i, deck := range m.decks {
    isRowSelected := i == m.hoveredDeck
    deckName := deck.name
    if isRowSelected {
      deckName = "→ " + deckName 
    }
    if i == len(m.decks) - 1 {
      rows = append(rows, styles.Row(true, isRowSelected, deckName, deck.review, deck.total, deck.createdAt)) 
    } else {
      rows = append(rows, styles.Row(false, isRowSelected, deckName, deck.review, deck.total, deck.createdAt))
    }
  }

	buttonStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#FF0000")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1)

	buttonStyle2 := buttonStyle.Background(lipgloss.Color("#4CAC00"))
	
  playButton := buttonStyle.MarginLeft(4).Render("▶ Play")
	addDeckButton := buttonStyle2.Margin(0, 8).Render("+ Add deck")
	addCardButton := buttonStyle2.MarginRight(4).Render("+ Add card")

	buttons := lipgloss.JoinHorizontal(0, playButton, addDeckButton, addCardButton)

  return lipgloss.JoinVertical(0, rows...) + "\n\n" + buttons
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
		}
	}
  return m, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error")
		os.Exit(1)
	}
}
