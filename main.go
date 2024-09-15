package main

import (
	"strconv"
	"strings"

	"fmt"
	"os"
	"time"

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
	PlayingDeckGuessing
	PlayingDeckResult
)

type model struct {
	decks            []Deck
	hoveredDeck      int
	selectedDeck     int
	selectedButton   int
	appState         AppState
	textInput        textinput.Model
	nameOfNewCard    string
	learnFlashcardID int
}

const layout = "2006-01-02"

func parseDate(dateStr string) time.Time {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return date
}

var uninitializedDecks = []Deck{
	{
		Name:      "Spanish 🇪🇸",
		CreatedAt: parseDate("2012-12-24"),
		Flashcards: []Flashcard{
			{FaceUp: "Hola", FaceDown: "Hello", IsLearned: false, ReviewDate: parseDate("2024-09-16")},
			{FaceUp: "Adiós", FaceDown: "Goodbye", IsLearned: true, ReviewDate: parseDate("2024-09-16")},
			{FaceUp: "Gracias", FaceDown: "Thank you", IsLearned: false, ReviewDate: parseDate("2024-09-17")},
			{FaceUp: "Por favor", FaceDown: "Please", IsLearned: true, ReviewDate: parseDate("2024-09-18")},
			{FaceUp: "Perdón", FaceDown: "Sorry", IsLearned: false, ReviewDate: parseDate("2024-09-19")},
		},
	},
	{
		Name:      "German 🇩🇪",
		CreatedAt: parseDate("2011-05-14"),
		Flashcards: []Flashcard{
			{FaceUp: "Hallo", FaceDown: "Hello", IsLearned: false, ReviewDate: parseDate("2024-09-20")},
			{FaceUp: "Tschüss", FaceDown: "Goodbye", IsLearned: true, ReviewDate: parseDate("2024-09-21")},
			{FaceUp: "Danke", FaceDown: "Thank you", IsLearned: false, ReviewDate: parseDate("2024-09-09")},
			{FaceUp: "Bitte", FaceDown: "Please", IsLearned: true, ReviewDate: parseDate("2024-09-23")},
			{FaceUp: "Entschuldigung", FaceDown: "Sorry", IsLearned: false, ReviewDate: parseDate("2024-09-11")},
		},
	},
	{
		Name:      "French 🇫🇷",
		CreatedAt: parseDate("2009-07-18"),
		Flashcards: []Flashcard{
			{FaceUp: "Bonjour", FaceDown: "Hello", IsLearned: true, ReviewDate: parseDate("2024-09-13")},
			{FaceUp: "Merci", FaceDown: "Thank you", IsLearned: false, ReviewDate: parseDate("2024-09-12")},
			{FaceUp: "Pomme", FaceDown: "Apple", IsLearned: true, ReviewDate: parseDate("2024-09-25")},
			{FaceUp: "Chat", FaceDown: "Cat", IsLearned: false, ReviewDate: parseDate("2024-09-22")},
			{FaceUp: "Maison", FaceDown: "House", IsLearned: true, ReviewDate: parseDate("2024-09-30")},
		},
	},
}

func initializeDecks(decks []Deck) *[]Deck {
	for deckIndex, deck := range decks {
		decks[deckIndex].Total = strconv.Itoa(len(decks[deckIndex].Flashcards))
		readyForReview := 0
		for _, card := range deck.Flashcards {
			if card.ReviewDate.Before(time.Now()) {
				readyForReview += 1
			}
		}
		decks[deckIndex].Review = strconv.Itoa(readyForReview)
	}
	return &decks
}

func initialModel() model {
	initializedDecks := initializeDecks(uninitializedDecks)

	return model{
		decks:          *initializedDecks,
		hoveredDeck:    0,
		selectedDeck:   -1,
		selectedButton: -1,
		appState:       DeckSelection,
		textInput:      InitTextinput(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	switch m.appState {
	case AddingDeck:
		return AddDeckMenu(m.textInput)
	case AddingCardFaceUp:
		return AddCardMenu(m.textInput, m.decks[m.selectedDeck].Name, true)
	case AddingCardFaceDown:
		return AddCardMenu(m.textInput, m.decks[m.selectedDeck].Name, false)
	case PlayingDeckGuessing:
		flashcardInLearning := m.decks[m.selectedDeck].Flashcards[m.learnFlashcardID]
		return PlayDeckMenu(m.textInput, flashcardInLearning.FaceUp, true, false)
	case PlayingDeckResult:
		flashcardInLearning := m.decks[m.selectedDeck].Flashcards[m.learnFlashcardID]
		if strings.ToLower(flashcardInLearning.FaceDown) == strings.ToLower(m.textInput.Value()) {
			return PlayDeckMenu(m.textInput, flashcardInLearning.FaceUp, false, true)
		} else {
      m.textInput.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color("#FF0000"))
      m.textInput.SetValue(flashcardInLearning.FaceDown)
			return PlayDeckMenu(m.textInput, flashcardInLearning.FaceUp, false, false)
		}
	}

	header := Header(len(m.decks) == 0)
	rows := []string{header}

	for i, deck := range m.decks {
		deckState := DeckState{
			Deck:           deck,
			IsDeckHovered:  i == m.hoveredDeck,
			IsDeckSelected: i == m.selectedDeck,
			IsBottomRow:    i == len(m.decks)-1,
		}
		rows = append(rows, Row(deckState))
	}

	return lipgloss.JoinVertical(0, rows...) + "\n\n" + ButtonMenuComponent(m.selectedButton)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if m.appState == DeckSelection || m.appState == ButtonMenu {
			if msg.String() == "A" {
				m.selectedButton = 2
				m = selectButton(m)
				return m, nil
			}
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
				m = selectButton(m)
			}
		case AddingDeck:
			switch msg.String() {
			case "enter":
				inputValue := m.textInput.Value()

				if inputValue != "" {
					deckValue := Deck{
						Name:      inputValue,
						CreatedAt: time.Now().Truncate(24 * time.Hour),
						Review:    "0",
						Total:     "0",
					}
					m.decks = append(m.decks, deckValue)
				}
				m = returnToMainMenu(m)
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case AddingCardFaceUp:
			switch msg.String() {
			case "enter":
				inputValue := m.textInput.Value()
				if inputValue != "" {
					m.nameOfNewCard = inputValue
					m.textInput.Reset()
					m.appState = AddingCardFaceDown
				}
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case AddingCardFaceDown:
			switch msg.String() {
			case "enter":
				inputValue := m.textInput.Value()
				if inputValue != "" {
					card := Flashcard{
						FaceUp:   m.nameOfNewCard,
						FaceDown: inputValue,
					}
					selectedDeck := &m.decks[m.selectedDeck]
					selectedDeck.Flashcards = append(selectedDeck.Flashcards, card)
					currentTotal, err := strconv.Atoi(selectedDeck.Total)
					if err != nil {
						fmt.Println("Erorr converting string to int")
					}
					newTotal := currentTotal + 1
					newTotalString := strconv.Itoa(newTotal)
					selectedDeck.Total = newTotalString
					m.nameOfNewCard = ""
					m = returnToMainMenu(m)
				}
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case PlayingDeckGuessing:
			switch msg.String() {
			case "enter":
				m.appState = PlayingDeckResult
			case "esc":
				m = returnToMainMenu(m)
			}
			m.textInput, _ = m.textInput.Update(msg)
		case PlayingDeckResult:
			switch msg.String() {
      case "esc":
        m = returnToMainMenu(m)
			}
		}

	}
	return m, nil
}

func findReviewCardID(slice []Flashcard) (cardID int) {
	for i := range slice {
		if slice[i].ReviewDate.Before(time.Now()) {
			return i
		}
	}
	return -1
}

func selectButton(m model) model {
	switch m.selectedButton {
	case 0:
		id := findReviewCardID(m.decks[m.selectedDeck].Flashcards)
		if id == -1 {
			m = returnToMainMenu(m)
		} else {
			m.learnFlashcardID = id
			m.appState = PlayingDeckGuessing
		}
	case 1:
		m.appState = AddingCardFaceUp
	case 2:
		m.appState = AddingDeck
	}
	return m
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
