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

// Header Badge Style (soft reddish background with dark bold text matching wireframe)
var headerStyle = gloss.NewStyle().
	Background(gloss.Color("203")).
	Foreground(gloss.Color("0")).
	Bold(true).
	Padding(0, 1)

// Left Navigation Sidebar Styles
var navActiveStyle = gloss.NewStyle().
	Foreground(gloss.Color("205")).
	Bold(true)

var navInactiveStyle = gloss.NewStyle().
	Foreground(gloss.Color("250"))

// Metric Selection Box & Card Styles
var cardSelectedStyle = gloss.NewStyle().
	BorderStyle(gloss.ThickBorder()).
	BorderForeground(gloss.Color("205"))

var cardNormalStyle = gloss.NewStyle().
	BorderStyle(gloss.NormalBorder()).
	BorderForeground(gloss.Color("240"))

// Detail Section & Text Info Styles
var infoBoxStyle = gloss.NewStyle().
	BorderStyle(gloss.NormalBorder()).
	BorderForeground(gloss.Color("240"))

var infoTitleStyle = gloss.NewStyle().
	Foreground(gloss.Color("212")).
	Bold(true)

var infoLabelStyle = gloss.NewStyle().
	Foreground(gloss.Color("245")).
	Bold(true)

var infoValueStyle = gloss.NewStyle().
	Foreground(gloss.Color("15")).
	Bold(true)

