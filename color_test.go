package colorful

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSprint(t *testing.T) {
	cases := []struct {
		name            string
		expectedValue   string
		text            string
		color           string
		backgroundColor string
	}{
		{
			name:            "check with default color and background",
			expectedValue:   "hello" + ResetColor,
			text:            "hello",
			color:           DefaultColor,
			backgroundColor: DefaultBackgroundColor,
		}, {
			name:            "check with not default color and not default background",
			expectedValue:   RedColor + BlueBackgroundColor + "hello" + ResetColor,
			text:            "hello",
			color:           RedColor,
			backgroundColor: BlueBackgroundColor,
		},
		{
			name:            "check with default color and not default background",
			expectedValue:   DefaultColor + RedBackgroundColor + "hello" + ResetColor,
			text:            "hello",
			color:           DefaultColor,
			backgroundColor: RedBackgroundColor,
		},
		{
			name:            "check with not default color and default background",
			expectedValue:   RedColor + DefaultBackgroundColor + "hello" + ResetColor,
			text:            "hello",
			color:           RedColor,
			backgroundColor: DefaultBackgroundColor,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, WithColor(c.color).WithBackgroundColor(c.backgroundColor).Sprint(c.text))
		})
	}
}

func TestSprintln(t *testing.T) {
	cases := []struct {
		name            string
		expectedValue   string
		text            string
		color           string
		backgroundColor string
	}{
		{
			name:            "check with default color and background",
			expectedValue:   "hello" + ResetColor + "\n",
			text:            "hello",
			color:           DefaultColor,
			backgroundColor: DefaultBackgroundColor,
		}, {
			name:            "check with not default color and not default background",
			expectedValue:   RedColor + BlueBackgroundColor + "hello" + ResetColor + "\n",
			text:            "hello",
			color:           RedColor,
			backgroundColor: BlueBackgroundColor,
		},
		{
			name:            "check with default color and not default background",
			expectedValue:   DefaultColor + RedBackgroundColor + "hello" + ResetColor + "\n",
			text:            "hello",
			color:           DefaultColor,
			backgroundColor: RedBackgroundColor,
		},
		{
			name:            "check with not default color and default background",
			expectedValue:   RedColor + DefaultBackgroundColor + "hello" + ResetColor + "\n",
			text:            "hello",
			color:           RedColor,
			backgroundColor: DefaultBackgroundColor,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, WithColor(c.color).WithBackgroundColor(c.backgroundColor).Sprintln(c.text))
		})
	}
}

func TestSprintf(t *testing.T) {
	cases := []struct {
		name            string
		expectedValue   string
		format          string
		parameter       any
		color           string
		backgroundColor string
	}{
		{
			name:            "check with default color and background",
			expectedValue:   "hello" + ResetColor,
			format:          "%s",
			parameter:       "hello",
			color:           DefaultColor,
			backgroundColor: DefaultBackgroundColor,
		}, {
			name:            "check with not default color and not default background",
			expectedValue:   RedColor + BlueBackgroundColor + "2" + ResetColor,
			format:          "%d",
			parameter:       2,
			color:           RedColor,
			backgroundColor: BlueBackgroundColor,
		},
		{
			name:            "check with default color and not default background",
			expectedValue:   DefaultColor + RedBackgroundColor + "3.14" + ResetColor,
			format:          "%1.2f",
			parameter:       3.14,
			color:           DefaultColor,
			backgroundColor: RedBackgroundColor,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expectedValue, WithColor(c.color).WithBackgroundColor(c.backgroundColor).Sprintf(c.format, c.parameter))
		})
	}
}
