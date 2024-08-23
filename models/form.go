package models

import (
	html "eightgoodthings/services"
	"eightgoodthings/styles"
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type FormKeys struct {
	About key.Binding
	Form  key.Binding
	Write key.Binding
	Quit  key.Binding
}

func (k FormKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.About}
}

func (k FormKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Write, k.About},
	}
}

var FormKeyMap = FormKeys{
	About: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("ctrl+a", "about"),
	),
	Form: key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl+f", "form"),
	),
	Write: key.NewBinding(
		key.WithKeys("ctrl+w"),
		key.WithHelp("ctrl+w", "write"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+q", "ctrl+c"),
		key.WithHelp("ctrl+q", "quit"),
	),
}

type FormModel struct {
	state    string
	about    tea.Model
	form     *huh.Form
	keys     FormKeys
	help     help.Model
	question string
	category string
	answers  []string
	width    int
	height   int
}

var fieldNames = []string{
	"First", "Second", "Third", "Fourth",
	"Fifth", "Sixth", "Seventh", "Eighth",
}

func NewForm(question string, answers int) FormModel {
	fields := []huh.Field{huh.NewInput().Title("Category").Key("Category")}

	for _, name := range fieldNames {
		input := huh.NewInput().Title(name).Key(name)
		fields = append(fields, input)
	}

	form := huh.NewForm(
		huh.NewGroup(fields...).Title(question).WithHeight(3),
	).WithWidth(50)

	return FormModel{
		state:    "form",
		about:    NewAbout(),
		keys:     FormKeyMap,
		help:     help.New(),
		form:     form,
		question: question,
		category: "",
		answers:  make([]string, answers),
	}
}

func (m FormModel) WriteHTML() tea.Cmd {
	html.UpdateFiles(m.category, m.answers)

	return nil
}

func (m FormModel) Init() tea.Cmd {
	if m.form == nil {
		return nil
	}

	return m.form.Init()
}

func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if m.form != nil && m.state == "form" {
		f, cmd := m.form.Update(msg)
		m.form = f.(*huh.Form)
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, FormKeyMap.About):
			m.state = "about"
			m.about, cmd = m.about.Update(msg)
		case key.Matches(msg, FormKeyMap.Form):
			m.state = "form"
		case key.Matches(msg, FormKeyMap.Write):
			cmds = append(cmds, m.WriteHTML())
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, FormKeyMap.Quit):
			cmds = append(cmds, tea.Quit)
		default:
			log.Printf("Form.Update(%v)::msg.KeyMsg", msg)
		}
	default:
		log.Printf("Form.Update(%v)::msg.msg.(type)", msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m FormModel) View() string {
	log.Printf("Form.View() w: %d h: %d", m.width, m.height)

	var output string

	if m.state == "form" {
		if m.form == nil {
			output = "Starting..."
		}

		if m.form.State == huh.StateCompleted {
			m.help.ShowAll = true

			responses := []string{styles.TitleStyle("RESPONSES")}
			m.category = m.form.GetString("Category")
			for i, name := range fieldNames {
				m.answers[i] = m.form.GetString(name)

				responses = append(
					responses, fmt.Sprintf("%v: %v", name, m.answers[i]),
				)
			}

			output = lipgloss.JoinVertical(
				lipgloss.Left,
				lipgloss.JoinVertical(
					lipgloss.Left,
					responses...,
				),
				m.help.View(m.keys),
			)
		} else {
			output = lipgloss.JoinVertical(
				lipgloss.Left,
				styles.TitleStyle(m.question),
				m.form.View(),
				m.help.View(m.keys),
			)
		}
	}

	if m.state == "about" {
		output = lipgloss.JoinVertical(
			lipgloss.Left,
			m.about.View(),
		)
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			output,
		),
	)
}
