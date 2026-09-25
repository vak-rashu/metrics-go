package tui

import (
	gloss "charm.land/lipgloss/v2"
)

var defaultStyle = gloss.NewStyle().
	BorderStyle(gloss.NormalBorder()).
	BorderForeground(gloss.Color("63"))

var cpuStyle = gloss.NewStyle().
	Foreground(gloss.Color("3")) // yellow
var diskReadStyle = gloss.NewStyle().
	Foreground(gloss.Color("2")) // green
var diskWriteStyle = gloss.NewStyle().
	Foreground(gloss.Color("1")) // red
var netRXStyle = gloss.NewStyle().
	Foreground(gloss.Color("6")) // cyan
var netTXStyle = gloss.NewStyle().
	Foreground(gloss.Color("5")) // magenta
