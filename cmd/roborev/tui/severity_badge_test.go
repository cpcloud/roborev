package tui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDerefOrZero(t *testing.T) {
	assert.Equal(t, 0, derefOrZero(nil), "nil pointer should yield 0")
	v := 42
	assert.Equal(t, 42, derefOrZero(&v), "non-nil pointer should yield underlying value")
	zero := 0
	assert.Equal(t, 0, derefOrZero(&zero), "explicit 0 pointer should yield 0")
}

func TestRenderSeverityBadge(t *testing.T) {
	tests := []struct {
		name             string
		h, m, l          int
		wantContains     []string
		wantPlainTextLen int
	}{
		{
			name:             "all zero",
			wantContains:     []string{"H0", "M0", "L0"},
			wantPlainTextLen: 8, // "H0 M0 L0"
		},
		{
			name: "single digits",
			h:    3, m: 2, l: 5,
			wantContains:     []string{"H3", "M2", "L5"},
			wantPlainTextLen: 8,
		},
		{
			name: "double digits",
			h:    12, m: 3, l: 8,
			wantContains:     []string{"H12", "M3", "L8"},
			wantPlainTextLen: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := renderSeverityBadge(tt.h, tt.m, tt.l)
			plain := stripANSI(out)
			assert.Len(t, plain, tt.wantPlainTextLen, "plain text length: got %q", plain)
			for _, want := range tt.wantContains {
				assert.Contains(t, plain, want,
					"want %q in %q", want, plain)
			}
		})
	}
}
