package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/marcelohmdias/dotforge/cli/print"
)

const banner = `
     __     __  ___
 ___/ /__  / /_/ _/__  _______  ___
/ _  / _ \/ __/ _/ _ \/ __/ _ \/ -_\
\_,_/\___/\__/_/ \___/_/  \_, /\__/
                         /___/
`

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("4"))

func Banner() {
	print.Println(style.Render(banner))
}
