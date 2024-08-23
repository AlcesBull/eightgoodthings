package models

import (
	"eightgoodthings/styles"
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AboutKeys struct {
	Form key.Binding
	Quit key.Binding
}

func (k AboutKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Form}
}

func (k AboutKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Form},
	}
}

var AboutKeyMap = AboutKeys{
	Form: key.NewBinding(
		key.WithKeys("alt+f"),
		key.WithHelp("alt+f", "form"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
}

type AboutModel struct {
	keys AboutKeys
	help help.Model
}

func NewAbout() AboutModel {
	return AboutModel{
		keys: AboutKeyMap,
		help: help.New(),
	}
}

func (m AboutModel) Init() tea.Cmd {
	return nil
}

func (m AboutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	default:
		log.Printf("AboutModel.Update(%v)::msg.(type)\n", msg)
	}

	return m, nil
}

func (m AboutModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.SectionStyle("8 Good Things"),
		styles.SectionStyle("created with 💘 using charm.io"),
		m.help.View(m.keys),
	)
}
