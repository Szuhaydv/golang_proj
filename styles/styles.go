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

var emptyHeaderDivider = lipgloss.NewStyle().
	Border(lipgloss.Border{
		Top:    "┬",
		Bottom: "┴",
	}, true, false).Render("│")

func EmptyHeader() string {
	leftCellBorder, middleCellBorder, rightCellBorder := defaultBorder, defaultBorder, defaultBorder

	leftCellBorder.TopLeft = "╭"
  leftCellBorder.BottomLeft = "╰"
	rightCellBorder.TopRight = "╮"
  rightCellBorder.BottomRight = "╯"

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

	return lipgloss.JoinHorizontal(0, leftCellStyle.Render("Decks"), emptyHeaderDivider, middleCellStyle.Render("Review / Total"), emptyHeaderDivider, rightCellStyle.Render("Created at"))
}
