package tui

import (
	"fmt"
	"strings"
)

// derefOrZero returns the dereferenced int, or 0 if the pointer is nil.
// Used to safely combine the three finding-count pointers that JSON
// deserialization or ad-hoc producers may populate inconsistently.
func derefOrZero(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

// renderSeverityBadge formats finding counts as "H3 M2 L5" with severity-
// colored letters/numbers. Zero counts render in dim grey so columns stay
// visually stable but de-emphasised. Plain-text width is always at least
// 8 characters (single digits) and grows by one per extra digit per slot.
func renderSeverityBadge(h, m, l int) string {
	parts := []string{
		formatSeveritySlot("H", h, severityHighStyle),
		formatSeveritySlot("M", m, severityMediumStyle),
		formatSeveritySlot("L", l, severityLowStyle),
	}
	return strings.Join(parts, " ")
}

// lipglossStyleRenderer is the minimal interface from lipgloss.Style that
// formatSeveritySlot consumes. lipgloss.Style satisfies it natively.
type lipglossStyleRenderer interface {
	Render(strs ...string) string
}

func formatSeveritySlot(letter string, count int, active lipglossStyleRenderer) string {
	text := fmt.Sprintf("%s%d", letter, count)
	if count == 0 {
		return severityZeroStyle.Render(text)
	}
	return active.Render(text)
}
