package command_test

import (
	"go_base/command"
	"testing"
)

func TestCommand(t *testing.T) {
	var tests = []struct {
		output int
	}{
		{ command.SUCCESS },
	}
	for _, test := range tests {
		result := command.Execute()
		if result != test.output {
			t.Errorf("Bad Exit: \"command.Execute()\" = %q not %q", result, test.output)
		}
	}
}
