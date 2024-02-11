package config

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// GetLogLevel
func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		name       string
		inputLevel string
		expected   zerolog.Level
	}{
		{name: "panic", inputLevel: "panic", expected: zerolog.PanicLevel},
		{name: "fatal", inputLevel: "fatal", expected: zerolog.FatalLevel},
		{name: "error", inputLevel: "error", expected: zerolog.ErrorLevel},
		{name: "warn", inputLevel: "warn", expected: zerolog.WarnLevel},
		{name: "info", inputLevel: "info", expected: zerolog.InfoLevel},
		{name: "debug", inputLevel: "debug", expected: zerolog.DebugLevel},
		{name: "trace", inputLevel: "trace", expected: zerolog.TraceLevel},
		{name: "empty", inputLevel: "", expected: zerolog.DebugLevel},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := test.inputLevel
			expected := test.expected

			actual := GetLogLevel(input)

			require.Equal(t, expected, actual)
		})
	}
}
