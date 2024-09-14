package styles

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type Flashcard struct {
  FaceUp string
  FaceDown string
  IsLearned bool
  ReviewDate time.Time
}

type Deck struct {
	Name      string
	Review    string
	Total     string
	CreatedAt time.Time
  Flashcards []Flashcard
}

type DeckState struct {
	Deck           Deck
	IsDeckHovered  bool
	IsDeckSelected bool
	IsBottomRow    bool
}

type SubMenu int

const (
  PlayDeck SubMenu = iota
  AddCard
  AddDeck
)

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

	return lipgloss.JoinHorizontal(0, leftCellStyle.Render(state.Deck.Name), cellDivider, middleCellStyle.Render(state.Deck.Review+" / "+state.Deck.Total), cellDivider, rightCellStyle.Render(state.Deck.CreatedAt.Format("2006-01-02")))
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
	addCardButton := buttonStyle2.Margin(0, checkIfButtonSelected(selectedButton, 2), 0, checkIfButtonSelected(selectedButton, 1)).Render("+ Add card")
	addDeckButton := buttonStyle2.Render("(A) Add deck")

	buttons := []string{playButton, addCardButton, addDeckButton}
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

func InitTextinput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter deck name"
	ti.Focus()
	ti.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#4CAC00")).Foreground(lipgloss.Color("#000000"))
	ti.CharLimit = 20
	ti.Width = 24
	return ti
}

func AddCardMenu(ti textinput.Model, deckName string, isFaceUp bool) string {
  return addingComponent(ti, deckName, isFaceUp, "")
}

func AddDeckMenu(ti textinput.Model) string {
  return addingComponent(ti, "", false, "")
}

func PlayDeckMenu(ti textinput.Model, word string) string {
  return addingComponent(ti, "", false, word)
}

func addingComponent(ti textinput.Model, deckName string, isFaceUp bool, word string) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(60)

	escText := lipgloss.NewStyle().Faint(true).Render("Esc")
	title := "Creating new deck"
  if deckName != "" {
    title = fmt.Sprintf("Adding new card to '%v'", deckName)
  } else if word != "" {
    labelText := lipgloss.NewStyle().MarginRight(4).Render("Word:")
    wordText := lipgloss.NewStyle().Foreground(lipgloss.Color("#FB9700")).Render(word)
    title = lipgloss.JoinHorizontal(0, labelText, wordText)
  }
	contentWidth := 60
	escWidth := lipgloss.Width(escText)
	titleWidth := lipgloss.Width(title)
	paddingLeft := (contentWidth - escWidth - titleWidth) / 2

	titleText := lipgloss.NewStyle().MarginLeft(paddingLeft).Bold(true).Render(title)

	titleRow := lipgloss.JoinHorizontal(lipgloss.Left, escText, titleText)

  textInputWidth := lipgloss.Width(ti.View())

	inputLabelWidth := escWidth + (contentWidth-escWidth-textInputWidth)/2
	labelStyle := lipgloss.NewStyle().
		Width(inputLabelWidth).
		Align(lipgloss.Right).
		MarginRight(2)

  label := "Name:"
  if deckName != "" {
    if isFaceUp {
      label = "Face up:"
    } else {
      label = "Face down:"
    }
  } else if word != "" {
    label = "Meaning:"
  }
	inputLabel := labelStyle.Render(label)

	inputRow := lipgloss.JoinHorizontal(lipgloss.Center, inputLabel, ti.View())

	return boxStyle.Render(lipgloss.JoinVertical(0, titleRow+"\n\n", inputRow+"\n\n"))

}

