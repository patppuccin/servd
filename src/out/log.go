package out

import (
	"fmt"
	"servd/src/constants"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleIcon        = lipgloss.NewStyle().Bold(true)
	styleMessage     = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))             // White/Grey
	styleBannerTitle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12")) // Blue
	styleBannerBlock = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))             // Dim gray

	iconDebug   = styleIcon.Foreground(lipgloss.Color("6")).Render("[?]")  // Cyan
	iconInfo    = styleIcon.Foreground(lipgloss.Color("12")).Render("[i]") // Blue
	iconWarn    = styleIcon.Foreground(lipgloss.Color("11")).Render("[!]") // Yellow
	iconError   = styleIcon.Foreground(lipgloss.Color("9")).Render("[x]")  // Red
	iconSuccess = styleIcon.Foreground(lipgloss.Color("10")).Render("[✓]") // Green
)

func Debug(msg string) {
	fmt.Printf("%s %s\n", iconDebug, styleMessage.Render(msg))
}

func Info(msg string) {
	fmt.Printf("%s %s\n", iconInfo, styleMessage.Render(msg))
}

func Warn(msg string) {
	fmt.Printf("%s %s\n", iconWarn, styleMessage.Render(msg))
}

func Error(msg string) {
	fmt.Printf("%s %s\n", iconError, styleMessage.Render(msg))
}

func Success(msg string) {
	fmt.Printf("%s %s\n", iconSuccess, styleMessage.Render(msg))
}

func Banner(message string) string {
	var sb strings.Builder

	sb.WriteString(styleBannerBlock.Render(constants.AppBanner))
	sb.WriteString("\n")

	sb.WriteString(styleBannerTitle.Render(message))
	sb.WriteString("\n")

	return sb.String()
}
