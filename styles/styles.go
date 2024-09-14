package styles

import (
	"github.com/charmbracelet/lipgloss"
  "github.com/charmbracelet/bubbles/textinput"
)

type Deck struct {
	Name      string
	Review    string
	Total     string
	CreatedAt string
}

type DeckState struct {
	Deck           Deck
	IsDeckHovered  bool
	IsDeckSelected bool
	IsBottomRow    bool
}

var defaultBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "┌",
	TopRight:    "┐",
	BottomLeft:  "├",
	BottomRight: "┘",
}

func Header(empty bool) string {

	headerDivider := lipgloss.NewStyle().
		Border(lipgloss.Border{
			Top: "┬",
			Bottom: func() string {
				if empty {
					return "┴"
				}
				return "┼"
			}(),
		}, true, false).Render("│")

	leftCellBorder, rightCellBorder := defaultBorder, defaultBorder

	leftCellBorder.TopLeft = "╭"
	rightCellBorder.TopRight = "╮"
	if empty {
		leftCellBorder.BottomLeft = "╰"
		rightCellBorder.BottomRight = "╯"
	} else {
		leftCellBorder.BottomLeft = "├"
		rightCellBorder.BottomRight = "┤"
	}

  headerCellStyle := lipgloss.NewStyle().PaddingLeft(2).Bold(true)

	leftCellStyle := headerCellStyle.
		Border(leftCellBorder, true, false, true, true).
		Width(24)

	middleCellStyle := headerCellStyle.
		Border(defaultBorder, true, false).
		Width(18)

	rightCellStyle := headerCellStyle.
		Border(rightCellBorder, true, true, true, false).
		Width(16)

	return lipgloss.JoinHorizontal(lipgloss.Bottom, 
    leftCellStyle.Render("Decks"), 
    headerDivider, 
    middleCellStyle.Render("Review / Total"), 
    headerDivider, 
    rightCellStyle.Render("Created at"),
  )
}

func Row(state DeckState) string {
	cellDivider := lipgloss.NewStyle().Border(lipgloss.Border{
		Bottom: func() string {
			if state.IsBottomRow {
				return "┴"
			}
			return "┼"
		}(),
	}, false, false, true, false).Render("│")

	leftCellBorder, rightCellBorder := defaultBorder, defaultBorder

	if state.IsBottomRow {
		leftCellBorder.BottomLeft = "╰"
		rightCellBorder.BottomRight = "╯"
	} else {
		leftCellBorder.BottomLeft = "├"
		rightCellBorder.BottomRight = "┤"
	}

	leftCellStyle := lipgloss.NewStyle().
		Border(leftCellBorder, false, false, true, true).
		PaddingLeft(2).
		Width(24)

	if state.IsDeckHovered || state.IsDeckSelected {
		leftCellStyle = leftCellStyle.Foreground(lipgloss.Color("#4CAC00"))
		if state.IsDeckHovered {
			state.Deck.Name = "→ " + state.Deck.Name
		}
	}

	middleCellStyle := lipgloss.NewStyle().
		Border(defaultBorder, false, false, true, false).
		Align(lipgloss.Center).
		Width(18)

	rightCellStyle := lipgloss.NewStyle().
		Border(rightCellBorder, false, true, true, false).
		Align(lipgloss.Center).
		Width(16)

	return lipgloss.JoinHorizontal(0, leftCellStyle.Render(state.Deck.Name), cellDivider, middleCellStyle.Render(state.Deck.Review+" / "+state.Deck.Total), cellDivider, rightCellStyle.Render(state.Deck.CreatedAt))
}

func checkIfButtonSelected(selectedButton int, buttonNo int) int {
	if selectedButton == buttonNo {
		return 0
	} else {
    if buttonNo == 1 || buttonNo == 2 {
      return 8
    }
		return 4
	}
}

func ButtonMenu(selectedButton int) string {
	buttonStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#FF0000")).
		Foreground(lipgloss.Color("#FFFFFF")).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1)

	buttonStyle2 := buttonStyle.Background(lipgloss.Color("#4CAC00"))

	playButton := buttonStyle.MarginLeft(checkIfButtonSelected(selectedButton, 0)).Render("▶ Play")
	addDeckButton := buttonStyle2.Margin(0, checkIfButtonSelected(selectedButton, 2), 0, checkIfButtonSelected(selectedButton, 1)).Render("+ Add deck")
	addCardButton := buttonStyle2.Render("+ Add card")

	buttons := []string{playButton, addDeckButton, addCardButton}
	if selectedButton != -1 {
		buttons = append(buttons[:selectedButton+1], buttons[selectedButton:]...)
    arrowMargin := 6
    if selectedButton == 0 {
      arrowMargin = 2
    }
		buttons[selectedButton] = lipgloss.NewStyle().MarginLeft(arrowMargin).Render("→ ")
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, buttons...)
}

func AddDeckMenu() textinput.Model {
  ti := textinput.New()
	ti.Placeholder = "Enter deck name"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 24
  return ti
}
