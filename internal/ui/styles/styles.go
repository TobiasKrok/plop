package styles

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

var (
	Colors = struct {
		Success color.Color
		Error   color.Color
	}{
		Success: lipgloss.Color("#8bbb11"),
		Error:   lipgloss.Color("#d32029"),
	}
)

func C(c color.Color, s string) string {
	return lipgloss.NewStyle().Foreground(c).Render(s)
}

func B(s string) string {
	return lipgloss.NewStyle().Bold(true).Render(s)
}

func BC(c color.Color, s string) string {
	return lipgloss.NewStyle().Foreground(c).Bold(true).Render(s)
}

func U(s string) string {
	return lipgloss.NewStyle().Underline(true).Render(s)
}

func UC(c color.Color, s string) string {
	return lipgloss.NewStyle().Foreground(c).Underline(true).Render(s)
}
