// Package cmd is here to provide a way for xk6 to build a binary with added extensions
package cmd

import (
	"context"

	"github.com/shiqinzhi/k6/cmd/state"
	internalcmd "github.com/shiqinzhi/k6/internal/cmd"
)

// Execute executes the k6 command
// It only is exported here for backwards compatibility and the ability to use xk6 to build extended k6
func Execute() {
	gs := state.NewGlobalState(context.Background())

	internalcmd.ExecuteWithGlobalState(gs)
}
