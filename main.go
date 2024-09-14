package main

import (
	"Szuhaydv/golang_proj/styles"

	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/lipgloss"
)

type AppState int

const (
	DeckSelection AppState = iota
	ButtonMenu
	AddingDeck
	AddingCardFaceUp
	AddingCardFaceDown
	PlayingDeck
)

type model struct {
	decks          []styles.Deck
	hoveredDeck    int
	selectedDeck   int
	selectedButton int
	appState       AppState
	textInput      textinput.Model
}

func initialModel() model {
	return model{
		decks: []styles.Deck{
			{
				Name:      "Spanish 🇪🇸",
				Review:    "15",
				Total:     "95",
				CreatedAt: "2012-12-24",
			},
			{
				Name:      "German 🇩🇪",
				Review:    "34",
				Total:     "128",
				CreatedAt: "2011-05-14",
			},
			{
				Name:      "English 🇬🇧",
				Review:    "9",
				Total:     "36",
				CreatedAt: "2009-07-18",
			},
		},
		hoveredDeck:    0,
		selectedDeck:   -1,
		selectedButton: -1,
		appState:       DeckSelection,
		textInput:      styles.InitTextinput(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	switch m.appState {
	case AddingDeck:
		return styles.AddDeckMenu(m.textInput)
	case AddingCardFaceUp:
		return styles.AddCardMenu(m.textInput, m.decks[m.selectedDeck].Name, true)
	case AddingCardFaceDown:
		return styles.AddCardMenu(m.textInput, m.decks[m.selectedDeck].Name, false)
	}
	header := styles.Header(len(m.decks) == 0)
	rows := []string{header}

	for i, deck := range m.decks {
		deckState := styles.DeckState{
			Deck:           deck,
			IsDeckHovered:  i == m.hoveredDeck,
			IsDeckSelected: i == m.selectedDeck,
			IsBottomRow:    i == len(m.decks)-1,
		}
		rows = append(rows, styles.Row(deckState))
	}

	return lipgloss.JoinVertical(0, rows...) + "\n\n" + styles.ButtonMenu(m.selectedButton)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		switch m.appState {
		case DeckSelection:
			switch msg.String() {
			case "up", "k":
				if m.hoveredDeck == -1 {
					break
				}
				if m.hoveredDeck > 0 {
					m.hoveredDeck--
				}

			case "down", "j":
				if m.hoveredDeck == -1 {
					break
				}
				if m.hoveredDeck < len(m.decks)-1 {
					m.hoveredDeck++
				}
			case "enter":
				m.selectedDeck = m.hoveredDeck
				m.hoveredDeck = -1
				m.selectedButton = 0
				m.appState = ButtonMenu
			}
		case ButtonMenu:
			switch msg.String() {

			case "right", "l":
				if m.selectedButton != -1 && m.selectedButton < 2 {
					m.selectedButton++
				}
			case "left", "h":
				if m.selectedButton != -1 && m.selectedButton > 0 {
					m.selectedButton--
				}
			case "enter":
				switch m.selectedButton {
				case 0:
					m.appState = PlayingDeck
				case 1:
					m.appState = AddingDeck
				case 2:
					m.appState = AddingCardFaceUp
				}
			}
		case AddingDeck:
			switch msg.String() {
			case "enter":
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case AddingCardFaceUp:
			switch msg.String() {
			case "enter":
				m.appState = AddingCardFaceDown
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case AddingCardFaceDown:
			switch msg.String() {
			case "enter":
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)

		}

	}
	return m, nil
}

func returnToMainMenu(m model) model {
	m.appState = DeckSelection
	m.hoveredDeck = 0
	m.selectedDeck = -1
	m.selectedButton = -1
	m.textInput.Reset()
	return m
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error")
		os.Exit(1)
	}
}
