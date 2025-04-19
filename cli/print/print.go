package print

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

var output = lipgloss.NewStyle().MarginLeft(1)

func Fprintln(w io.Writer, str string) {
	fmt.Fprintln(w, Render(str))
}

func Println(str string) {
	fmt.Println(Render(str))
}

func Render(str string) string {
	return output.Render(str)
}
