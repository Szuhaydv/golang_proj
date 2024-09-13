package main

import (
	"Szuhaydv/golang_proj/styles"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	// "os"
	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

// type Deck struct {
// 	name      string
// 	review    int
// 	total     int
// 	createdAt string
// }
//
// type model struct {
// 	decks       []Deck
// 	hoveredDeck int
// }

// func (m model) View() string {
//
// 	var myCuteBorder = lipgloss.Border{
// 		Top:         "─",
// 		Bottom:      "─",
// 		Left:        "│",
// 		Right:       "│",
// 		TopRight:    "╮",
// 		TopLeft:     "╭",
// 		BottomLeft:  "╰",
// 		BottomRight: "╯",
// 	}
//
// 	style := lipgloss.NewStyle().
// 		BorderStyle(myCuteBorder).
// 		BorderForeground(lipgloss.Color("205"))
//
// 	return style.Render("Hello Kitty")
// }

func main() {
	// p := tea.NewProgram(initialModel)
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Error")
	// 	os.Exit(1)
	// }

	buttonStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#FF0000")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Border(lipgloss.RoundedBorder()).
    Padding(0, 1)

  buttonStyle2 := buttonStyle.Background(lipgloss.Color("#4CAC00"))

	header := styles.Header(false)
	row1 := styles.Row(false, "Spanish 🇪🇸", "15", "95", "2012-12-14")
	row2 := styles.Row(true, "Spanish 🇪🇸", "15", "95", "2012-12-14")
	table := lipgloss.JoinVertical(0, header, row1, row2)
  playButton := buttonStyle.MarginLeft(4).Render("▶ Play")
  addDeckButton := buttonStyle2.Margin(0, 8).Render("+ Add deck")
  addCardButton := buttonStyle2.MarginRight(4).Render("+ Add card")

  buttons := lipgloss.JoinHorizontal(0, playButton, addDeckButton, addCardButton)

	fmt.Println(table + "\n\n" + buttons)
}
