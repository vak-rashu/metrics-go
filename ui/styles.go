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

// Services Tab Styles
var serviceTabFailedActiveStyle = gloss.NewStyle().
	Background(gloss.Color("203")).
	Foreground(gloss.Color("0")).
	Bold(true).
	Padding(0, 2)

var serviceTabRunningActiveStyle = gloss.NewStyle().
	Background(gloss.Color("2")).
	Foreground(gloss.Color("0")).
	Bold(true).
	Padding(0, 2)

var serviceTabDeadActiveStyle = gloss.NewStyle().
	Background(gloss.Color("242")).
	Foreground(gloss.Color("15")).
	Bold(true).
	Padding(0, 2)

var serviceTabInactiveStyle = gloss.NewStyle().
	Foreground(gloss.Color("250")).
	Padding(0, 2)

var serviceStatusFailedStyle = gloss.NewStyle().
	Foreground(gloss.Color("1")) // Red

var serviceStatusRunningStyle = gloss.NewStyle().
	Foreground(gloss.Color("2")) // Green

var serviceStatusDeadStyle = gloss.NewStyle().
	Foreground(gloss.Color("242")) // Gray

var serviceItemActiveStyle = gloss.NewStyle().
	Foreground(gloss.Color("205")).
	Bold(true)

var serviceItemInactiveStyle = gloss.NewStyle().
	Foreground(gloss.Color("252"))

var logHeaderStyle = gloss.NewStyle().
	Foreground(gloss.Color("212")).
	Bold(true)

var logEntryStyle = gloss.NewStyle().
	Foreground(gloss.Color("252"))


