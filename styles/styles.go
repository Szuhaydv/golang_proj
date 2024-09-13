package styles

import "github.com/charmbracelet/lipgloss"

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

	leftCellBorder, middleCellBorder, rightCellBorder := defaultBorder, defaultBorder, defaultBorder

	leftCellBorder.TopLeft = "╭"
	rightCellBorder.TopRight = "╮"
	if empty {
		leftCellBorder.BottomLeft = "╰"
		rightCellBorder.BottomRight = "╯"
	} else {
		leftCellBorder.BottomLeft = "├"
		rightCellBorder.BottomRight = "┤"
	}

	leftCellStyle := lipgloss.NewStyle().
		Border(leftCellBorder, true, false, true, true).
		PaddingLeft(2).
		Bold(true).
		Width(24)

	middleCellStyle := lipgloss.NewStyle().
		Border(middleCellBorder, true, false, true, false).
		PaddingLeft(2).
		Bold(true).
		Width(18)

	rightCellStyle := lipgloss.NewStyle().
		Border(rightCellBorder, true, true, true, false).
		PaddingLeft(2).
		Bold(true).
		Width(16)

	return lipgloss.JoinHorizontal(0, leftCellStyle.Render("Decks"), headerDivider, middleCellStyle.Render("Review / Total"), headerDivider, rightCellStyle.Render("Created at"))
}

func Row(bottom bool, deckName string, review string, total string, createdAt string) string {
  cellDivider := lipgloss.NewStyle().Border(lipgloss.Border{
			Bottom: func() string {
				if bottom {
					return "┴"
				}
				return "┼"
			}(),
  }, false, false, true, false).Render("│")

  leftCellBorder, rightCellBorder := defaultBorder, defaultBorder

	if bottom {
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

	middleCellStyle := lipgloss.NewStyle().
		Border(defaultBorder, false, false, true, false).
    Align(lipgloss.Center).
		Width(18)

	rightCellStyle := lipgloss.NewStyle().
		Border(rightCellBorder, false, true, true, false).
		Align(lipgloss.Center).
		Width(16)
  
  return lipgloss.JoinHorizontal(0, leftCellStyle.Render(deckName), cellDivider, middleCellStyle.Render(review + " / " + total), cellDivider, rightCellStyle.Render(createdAt))
}
