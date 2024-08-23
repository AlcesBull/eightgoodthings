package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle   = lipgloss.NewStyle().Bold(true).Height(2).Foreground(lipgloss.Color("190")).Render
	SectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("140")).Render
)
