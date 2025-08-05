// Package tests contains integration tests for multiple packages.
package tests

import (
	"testing"

	"github.com/shiqinzhi/k6/internal/cmd"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	Main(m)
}

func TestRootCommand(t *testing.T) {
	t.Parallel()

	cases := map[string][]string{
		"Just root": {"k6"},
		"Help flag": {"k6", "--help"},
	}

	helptxt := "Usage:\n  k6 [command]\n\nAvailable Commands"
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			ts := NewGlobalTestState(t)
			ts.CmdArgs = args
			cmd.ExecuteWithGlobalState(ts.GlobalState)
			assert.Len(t, ts.LoggerHook.Drain(), 0)
			assert.Contains(t, ts.Stdout.String(), helptxt)
		})
	}
}
