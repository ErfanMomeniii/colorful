package colorful

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSprint(t *testing.T) {
	cases := []struct {
		name          string
		expectedValue string
		text          string
		color         string
		background    string
	}{
		{
			name:          "check with default color and background",
			expectedValue: "hello" + ResetColor,
			text:          "hello",
			color:         DefaultColor,
			background:    DefaultBackground,
		}, {
			name:          "check with not default color and not default background",
			expectedValue: RedColor + BlueBackground + "hello" + ResetColor,
			text:          "hello",
			color:         RedColor,
			background:    BlueBackground,
		},
		{
			name:          "check with default color and not default background",
			expectedValue: DefaultColor + RedBackground + "hello" + ResetColor,
			text:          "hello",
			color:         DefaultColor,
			background:    RedBackground,
		},
		{
			name:          "check with not default color and default background",
			expectedValue: RedColor + DefaultBackground + "hello" + ResetColor,
			text:          "hello",
			color:         RedColor,
			background:    DefaultBackground,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, Sprint(c.color, c.background, c.text))
		})
	}
}

func TestSprintln(t *testing.T) {
	cases := []struct {
		name          string
		expectedValue string
		text          string
		color         string
		background    string
	}{
		{
			name:          "check with default color and background",
			expectedValue: "hello" + ResetColor + "\n",
			text:          "hello",
			color:         DefaultColor,
			background:    DefaultBackground,
		}, {
			name:          "check with not default color and not default background",
			expectedValue: RedColor + BlueBackground + "hello" + ResetColor + "\n",
			text:          "hello",
			color:         RedColor,
			background:    BlueBackground,
		},
		{
			name:          "check with default color and not default background",
			expectedValue: DefaultColor + RedBackground + "hello" + ResetColor + "\n",
			text:          "hello",
			color:         DefaultColor,
			background:    RedBackground,
		},
		{
			name:          "check with not default color and default background",
			expectedValue: RedColor + DefaultBackground + "hello" + ResetColor + "\n",
			text:          "hello",
			color:         RedColor,
			background:    DefaultBackground,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, Sprintln(c.color, c.background, c.text))
		})
	}
}

func TestSprintf(t *testing.T) {
	cases := []struct {
		name          string
		expectedValue string
		format        string
		parameter     any
		color         string
		background    string
	}{
		{
			name:          "check with default color and background",
			expectedValue: "hello" + ResetColor,
			format:        "%s",
			parameter:     "hello",
			color:         DefaultColor,
			background:    DefaultBackground,
		}, {
			name:          "check with not default color and not default background",
			expectedValue: RedColor + BlueBackground + "2" + ResetColor,
			format:        "%d",
			parameter:     2,
			color:         RedColor,
			background:    BlueBackground,
		},
		{
			name:          "check with default color and not default background",
			expectedValue: DefaultColor + RedBackground + "3.14" + ResetColor,
			format:        "%1.2f",
			parameter:     3.14,
			color:         DefaultColor,
			background:    RedBackground,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, Sprintf(c.color, c.background, c.format, c.parameter))
		})
	}
}
