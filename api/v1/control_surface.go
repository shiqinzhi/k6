package v1

import (
	"context"

	"github.com/shiqinzhi/k6/internal/execution"
	"github.com/shiqinzhi/k6/internal/metrics/engine"
	"github.com/shiqinzhi/k6/lib"
	"github.com/shiqinzhi/k6/metrics"
)

// ControlSurface includes the methods the REST API can use to control and
// communicate with the rest of k6.
type ControlSurface struct {
	RunCtx        context.Context
	Samples       chan metrics.SampleContainer
	MetricsEngine *engine.MetricsEngine
	Scheduler     *execution.Scheduler
	RunState      *lib.TestRunState
}
