package tui

import (
	gloss "charm.land/lipgloss/v2"
)

var headLineStyle = gloss.NewStyle().
	Width(100).
	Align(gloss.Center).
	Background(gloss.Color("3")).
	Foreground(gloss.Color("12"))

var tabStyle = gloss.NewStyle().
	Border(gloss.RoundedBorder()).
	PaddingLeft(1).
	PaddingRight(1).
	BorderForeground(gloss.Color("183"))

var blocks = []string{
	"CPU",
	"Processes",
	"Memory",
}
