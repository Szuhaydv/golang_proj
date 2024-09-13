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

  header := styles.Header(false)
  row := styles.Row(true, "Spanish 🇪🇸", "15", "95", "2012-12-14")
  table := lipgloss.JoinVertical(0, header, row)


  fmt.Println(table) 
}
